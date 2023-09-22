package main

import (
	"backend/pkg/database"
	"backend/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	db, err := database.Init()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	// Seed data
	err = database.SeedData(db, "/Users/aryaarun/Desktop/stripe-integration/backend/cmd/orders_sample.json")
	if err != nil {
		log.Fatalf("Error seeding data: %v", err)
	}
	http.HandleFunc("/orders", handlers.GetOrdersHandler(db))
	http.HandleFunc("/updateShipmentStatus", handlers.UpdateShipmentStatusHandler(db))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
