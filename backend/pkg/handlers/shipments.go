package handlers

import (
	"backend/pkg/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateShipmentStatusHandler(db *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			Status string `json:"status"`
		}
		err = json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid input in request body", http.StatusBadRequest)
			return
		}

		if err := db.Client().Ping(context.TODO(), nil); err != nil {
			log.Printf("Database connection error: %v", err)
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}

		shipmentsCollection := db.Collection("orders")
		filter := bson.M{"_id": objectID}
		log.Printf("Filter: %v", filter)

		var existingShipment models.Shipment
		err = shipmentsCollection.FindOne(context.TODO(), filter).Decode(&existingShipment)
		if err != nil {
			log.Printf("Error finding shipment: %v", err)
			http.Error(w, "Error finding shipment", http.StatusInternalServerError)
			return
		}

		if len(existingShipment.Fulfillments) > 0 {
			existingShipment.Fulfillments[0].ShipmentStatus = requestBody.Status
		}

		update := bson.M{"$set": bson.M{"fulfillments.0.shipment_status": requestBody.Status}}
		_, err = shipmentsCollection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Printf("Error updating shipment status: %v", err)
			http.Error(w, "Error updating shipment status", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Shipment status updated successfully"})
	}
}
