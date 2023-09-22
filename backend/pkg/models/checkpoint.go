package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Checkpoint struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ShipmentID     primitive.ObjectID `bson:"shipmentId" json:"-"`
	Status         string             `bson:"status"`
	Title          string             `bson:"title"`
	Timestamp      time.Time          `bson:"timestamp"`
	TrackerEventId string             `bson:"trackerEventId"`
	CreatedAt      time.Time          `bson:"createdAt"`
}
