package database

import (
	"backend/pkg/models"
	"context"
	"log"
	"strconv"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	stripe.Key = "sk_test_51L31C6SCAE3bkRmmEEsNiKDTAXjbWKapCWuUVNjYtdPHHJQUHJFZCsXSpMvDTCpfjEEGLnhjTk86wwikIKU5qIxP00mn5Nzx2A"
}

func SeedDataFromMongoToStripe(mongoURI string, dbName string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	orderCollection := client.Database(dbName).Collection("orders")
	cursor, err := orderCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	var orders []models.Order
	if err = cursor.All(context.TODO(), &orders); err != nil {
		return err
	}

	for _, order := range orders {
		// Create Customer in Stripe
		customerParams := &stripe.CustomerParams{
			Email: stripe.String(order.CustomerEmail),
			Name:  stripe.String(order.CustomerName.FirstName + " " + order.CustomerName.LastName),
			Phone: stripe.String(order.CustomerPhone),
			// Address: &stripe.AddressParams{
			// 	Line1:      stripe.String(order.CustomerAddress),
			// 	City:       stripe.String(order.ShippingAddress.City),
			// 	State:      stripe.String(order.ShippingAddress.Province),
			// 	PostalCode: stripe.String(order.ShippingAddress.Zip),
			// 	Country:    stripe.String(order.ShippingAddress.CountryCode),
			// }
		}
		c, err := customer.New(customerParams)
		if err != nil {
			log.Printf("Error creating customer for order %s: %v", order.CustomerEmail, err)
			continue
		}

		totalFloat, err := strconv.ParseFloat(order.OrderTotal, 64)
		if err != nil {
			log.Printf("Error parsing order total for order %s: %v", order.CustomerEmail, err)
			continue
		}
		amount := int64(totalFloat * 100)

		paymentIntentParams := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(amount),
			Currency: stripe.String(order.CustomerName.Currency),
			Customer: stripe.String(c.ID),
		}
		paymentIntent, err := paymentintent.New(paymentIntentParams)
		if err != nil {
			log.Printf("Error creating payment intent for order %s: %v", order.CustomerEmail, err)
		} else {
			// Update the order with the Stripe payment ID
			orderUpdate := bson.M{"$set": bson.M{"stripe_payment_id": paymentIntent.ID}}
			_, err = orderCollection.UpdateOne(context.TODO(), bson.M{"_id": order.ID}, orderUpdate)
			if err != nil {
				log.Printf("Error updating order with Stripe payment ID: %v", err)
			}
		}

	}
	return nil
}
