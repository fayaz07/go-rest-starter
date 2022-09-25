package db

import (
	"context"
	log "go-rest-starter/src/core/logger"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var deviceCollection *mongo.Collection

const DEVICE_COLLECTION = "devices"

type DeviceModel struct {
	ID            primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	SessionID     primitive.ObjectID `json:"sId" binding:"required" bson:"_sId"`
	UserID        primitive.ObjectID `json:"uId" bson:"_uId"`
	Type          string             `json:"type" binding:"required" bson:"type"`
	Connection    string             `json:"conn" binding:"required" bson:"conn"`
	OS            string             `json:"os" binding:"required" bson:"os"`
	OSVersion     string             `json:"osV" binding:"required" bson:"osV"`
	ClientVersion string             `json:"cV" binding:"required" bson:"cV"`

	Maker string `json:"maker" binding:"required" bson:"maker"`
	Model string `json:"model" binding:"required" bson:"model"`

	UserAgent        string `json:"ua" bson:"ua"`
	UserAgentVersion string `json:"uaV" bson:"uaV"`
	Referer          string `json:"referer" bson:"referer"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
}

func GetDeviceCollection(db *mongo.Database) *mongo.Collection {
	var once sync.Once
	once.Do(func() {
		log.I("trying to get database collection")
		deviceCollection = db.Collection(DEVICE_COLLECTION)
		createDeviceCollectionIndexes()
	})
	return deviceCollection
}

func createDeviceCollectionIndexes() {
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "ip", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("unique_fields_client"),
		},
	}
	opts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	_, err := deviceCollection.Indexes().CreateMany(context.Background(),
		models, opts)
	if err != nil {
		log.F(err.Error())
	}
}
