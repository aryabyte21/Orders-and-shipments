package handlers

import (
	"backend/pkg/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	stripe.Key = "sk_test_51L31C6SCAE3bkRmmEEsNiKDTAXjbWKapCWuUVNjYtdPHHJQUHJFZCsXSpMvDTCpfjEEGLnhjTk86wwikIKU5qIxP00mn5Nzx2A"
}

func UpdatePaymentMetadata(stripePaymentID string, trackingNumber string) {
	params := &stripe.PaymentIntentParams{}
	params.Params.Metadata = map[string]string{
		"tracking_number": trackingNumber,
	}

	_, err := paymentintent.Update(stripePaymentID, params)
	if err != nil {
		log.Printf("Error updating payment metadata: %v", err)
	}
}
func UpdateShipmentStatusHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		shipmentID := r.URL.Query().Get("shipmentID")
		if shipmentID == "" {
			http.Error(w, "Shipment ID not provided in URL parameters", http.StatusBadRequest)
			return
		}

		objectID, err := primitive.ObjectIDFromHex(shipmentID)
		if err != nil {
			http.Error(w, "Invalid shipment ID format", http.StatusBadRequest)
			return
		}

		var requestBody struct {
			Status         string `json:"status"`
			TrackingNumber string `json:"trackingNumber"`
		}
		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid input in request body", http.StatusBadRequest)
			return
		}

		shipmentsCollection := db.Collection("orders")
		filter := bson.M{"_id": objectID}

		var existingShipment models.Shipment
		err = shipmentsCollection.FindOne(context.TODO(), filter).Decode(&existingShipment)
		if err != nil {
			http.Error(w, "Error finding shipment", http.StatusInternalServerError)
			return
		}

		if len(existingShipment.Fulfillments) > 0 {
			existingShipment.Fulfillments[0].ShipmentStatus = requestBody.Status
		}

		update := bson.M{"$set": bson.M{"fulfillments.0.shipment_status": requestBody.Status}}
		_, err = shipmentsCollection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			http.Error(w, "Error updating shipment status", http.StatusInternalServerError)
			return
		}

		// Sync the tracking number with Stripe if the shipment status is "shipped" and we have a Stripe payment ID
		if existingShipment.StripePaymentID != "" && requestBody.Status == "shipped" {
			UpdatePaymentMetadata(existingShipment.StripePaymentID, requestBody.TrackingNumber)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Shipment status and tracking number updated successfully"})
	}
}
