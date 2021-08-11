package core

import (
	"fmt"
	"log"

	helpers "go-rest-starter/src/utils/helpers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	localhost        = "localhost"
	mongodbLocal     = "mongodb://%s:%s"
	mongodbLocalAuth = "mongodb://%s:%s@%s:%s"
	mongodbSRV       = "mongodb+srv://%s:%s@%s"
)

var didInitializeDb bool = false
var db *mongo.Database

// GetDbConnection (config DatabaseConfig) ... This will export database connection
func GetDbConnection(config *DatabaseConfig) *mongo.Database {
	if didInitializeDb {
		return db
	}
	connectToDb(config)
	return db
}

func connectToDb(config *DatabaseConfig) {

	connectionURI := mongodbConnString(config)
	// log.Println(connectionURI)

	ctx := helpers.GetContext()

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create MongoDb client: %v", err)
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to MongoDb cluster: %v", err)
		panic(err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
		panic(err)
	}

	db = client.Database(config.DbName)
	didInitializeDb = true
	log.Println("Connected to MongoDB!")
}

func mongodbConnString(config *DatabaseConfig) string {
	// for localhost: mongodb://<username>:<password>@host/db-name
	// only localhost will have a port
	// localhost dbs will not include srv
	if config.DbHost == localhost {

		if len(config.DbUsername) == 0 || len(config.DbPassword) == 0 {
			return fmt.Sprintf(mongodbLocal, config.DbHost, config.DbPort)
		}
		return fmt.Sprintf(mongodbLocalAuth, config.DbUsername, config.DbPassword, config.DbHost, config.DbPort)
	}

	// for remote host: mongodb+srv://<username>:<password>@host/db-name
	// atlas clusters will not provide a port or connect to default port 27017
	return fmt.Sprintf(mongodbSRV, config.DbUsername, config.DbPassword, config.DbHost)
}
