package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shipment struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	OrderID                 primitive.ObjectID `bson:"orderId" json:"-"`
	BusinessID              primitive.ObjectID `bson:"businessId" json:"-"`
	Status                  string             `bson:"status"`
	ProviderFulfillmentId   string             `bson:"providerFulfillmentId"`
	CarrierName             string             `bson:"carrierName"`
	CarrierCode             string             `bson:"carrierCode"`
	CarrierTrackingNumber   string             `bson:"carrierTrackingNumber"`
	EstimatedDeliveryDate   time.Time          `bson:"estimatedDeliveryDate"`
	ActualDeliveryDate      time.Time          `bson:"actualDeliveryDate"`
	CreatedAt               time.Time          `bson:"createdAt"`
	ProviderCreatedAt       time.Time          `bson:"providerCreatedAt"`
	Summary                 string             `bson:"summary"`
	DestinationZipCode      string             `bson:"destinationZipCode"`
	DestinationCity         string             `bson:"destinationCity"`
	DestinationState        string             `bson:"destinationState"`
	DestinationCountryCode  string             `bson:"destinationCountryCode"`
	DestinationLatitude     float64            `bson:"destinationLatitude"`
	DestinationLongitude    float64            `bson:"destinationLongitude"`
	OriginZipCode           string             `bson:"originZipCode"`
	OriginCity              string             `bson:"originCity"`
	OriginState             string             `bson:"originState"`
	OriginCountryCode       string             `bson:"originCountryCode"`
	ReturnShipmentID        primitive.ObjectID `bson:"returnShipmentId" json:"-"`
	Weight                  string             `bson:"weight"`
	TrackerID               string             `bson:"trackerId"`
	ProviderTrackingPageURL string             `bson:"providerTrackingPageUrl"`
	TrackerServiceCode      string             `bson:"trackerServiceCode"`
	UniqueId                string             `bson:"uniqueId"`
	Fulfillments            []Fulfillment      `bson:"fulfillments" json:"fulfillments"`
	StripePaymentID         string             `bson:"stripe_payment_id,omitempty" json:"stripe_payment_id,omitempty"`
}

type Fulfillment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ShipmentStatus string             `bson:"shipment_status" json:"shipment_status"`
}
