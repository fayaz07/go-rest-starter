package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionModel struct {
	ClientId  primitive.ObjectID `json:"clientId"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	AuthType  string             `json:"authType" bson:"authType"`
}
