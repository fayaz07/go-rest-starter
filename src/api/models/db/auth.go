package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var authCollection *mongo.Collection

const AUTH_COLLECTION = "auths"

type AuthModel struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`

	Email    string               `json:"email" binding:"required" bson:"email"`
	Password string               `json:"password" binding:"required" bson:"password"`
	Role     AuthRole             `json:"role" binding:"required" bson:"role"`
	Provider AuthProvider         `json:"provider" binding:"required" bson:"provider"`
	Sessions []primitive.ObjectID `json:"sessions" bson:"sessions"`
}
