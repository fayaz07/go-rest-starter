package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionModel struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`

	AuthType string       `json:"authType" bson:"authType"`
	IP       string       `json:"ip" bson:"ip"`
	DeviceID string       `json:"devId" bson:"devId"`
	Events   []EventModel `json:"events" bson:"events"`
}
