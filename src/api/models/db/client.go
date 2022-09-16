package db

import (
	"context"
	"sync"
	"time"

	log "go-rest-starter/src/core/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientCollection *mongo.Collection

const CLIENT_COLLECTION = "clients"

type ClientModel struct {
	ID        primitive.ObjectID `json:"id" binding:"required" bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`

	IP      string `json:"ip" binding:"required" bson:"ip"`
	Device  string `json:"device" binding:"required" bson:"device"`
	OS      string `json:"os" binding:"required" bson:"os"`
	Version string `json:"version" bson:"version"`

	// Fields for mobile
	Model        string `json:"model" bson:"model"`
	Manufacturer string `json:"manufacturer" bson:"manufacturer"`
	Connection   string `json:"connection" bson:"connection"`

	// Fields for web
	UserAgent string `json:"userAgent" bson:"userAgent"`
	Host      string `json:"host" bson:"host"`
	Browser   string `json:"browser" bson:"browser"`

	Referer string `json:"referer" bson:"referer"`
}

func GetClientCollection(db *mongo.Database) *mongo.Collection {
	var once sync.Once
	once.Do(func() {
		log.I("trying to get database collection")
		clientCollection = db.Collection(CLIENT_COLLECTION)
		// if clientCollection == nil {
		// 	// log.Fatal("clientCollection is nil")
		// 	db.CreateCollection(context.Background(), CLIENT_COLLECTION)
		// 	clientCollection = db.Collection(CLIENT_COLLECTION)
		// }

		createClientCollectionIndexes()
	})
	// clientCollection.InsertOne(context.Background(), bson.M{"foo": "bar"})
	return clientCollection
}

func createClientCollectionIndexes() {
	models := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "ip", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("unique_fields_client"),
		},
	}
	opts := options.CreateIndexes().SetMaxTime(2 * time.Second)
	_, err := clientCollection.Indexes().CreateMany(context.Background(),
		models, opts)

	if err != nil {
		log.F(err.Error())
	}
}

func (m ClientModel) Print() {
	PrettyPrint(m)
}
