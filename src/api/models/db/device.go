package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MOBILE_CLIENT  = "mobile"
	DESKTOP_CLIENT = "desktop"
	WEB_CLIENT     = "web"
)

type DeviceModel struct {
	ID            primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"uId" binding:"required" bson:"_uId"`
	Type          string             `json:"type" binding:"required" bson:"type"`
	OS            string             `json:"os" binding:"required" bson:"os"`
	OSVersion     string             `json:"osV" binding:"required" bson:"osV"`
	ClientVersion string             `json:"cV" binding:"required" bson:"cV"`

	Maker string `json:"maker" binding:"required" bson:"maker"`
	Model string `json:"model" binding:"required" bson:"model"`

	UserAgent        string `json:"ua" bson:"ua"`
	UserAgentVersion string `json:"uaV" bson:"uaV"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
}
