package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminCollection *mongo.Collection

const ADMIN_COLLECTION = "admins"

type AdminModel struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`

	Email     string `json:"email" binding:"required" bson:"email"`
	Password  string `json:"pwd" binding:"required" bson:"pwd"`
	FirstName string `json:"firstName" binding:"required" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}
