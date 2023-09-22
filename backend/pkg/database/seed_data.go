package database

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func SeedData(db *mongo.Database, filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}


	orders, ok := jsonData["orders"].([]interface{})
	if ok {
		ordersCollection := db.Collection("orders")
		_, err = ordersCollection.InsertMany(context.TODO(), orders)
		if err != nil {
			return err
		}
		log.Println("Seeded orders into MongoDB.")
	}

	return nil
}
