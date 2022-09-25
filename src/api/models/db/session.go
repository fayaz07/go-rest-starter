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

var sessionCollection *mongo.Collection

const SESSION_COLLECTION = "sessions"

type SessionModel struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`

	AuthType string       `json:"authType" bson:"authType"`
	IP       string       `json:"ip" binding:"required" bson:"ip"`
	Events   []EventModel `json:"events" bson:"events"`
}

func GetSessionCollection(db *mongo.Database) *mongo.Collection {
	var once sync.Once
	once.Do(func() {
		log.I("trying to get session collection")
		sessionCollection = db.Collection(SESSION_COLLECTION)
		createSessionCollectionIndexes()
	})
	return deviceCollection
}

func createSessionCollectionIndexes() {
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "ip", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("unique_fields_client"),
		},
	}
	opts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	_, err := sessionCollection.Indexes().CreateMany(context.Background(), models, opts)
	if err != nil {
		log.F(err.Error())
	}
}
