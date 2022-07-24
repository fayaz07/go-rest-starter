package core

import (
	"fmt"
	"log"
	"sync"

	"go-rest-starter/src/api/repository"
	appTypes "go-rest-starter/src/core/types"
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

var isDbConnectionAlive bool = false
var _db *mongo.Database
var _onceForDb sync.Once

func isDbConnectionHealthy() bool {
	err := _db.Client().Ping(ctx, nil)
	if err != nil {
		isDbConnectionAlive = false
	}
	return isDbConnectionAlive
}

func GetDbConnection() *mongo.Database {
	return _db
}

func InitializeDatabseConn(config *appTypes.DatabaseConfig) {
	_onceForDb.Do(func() {
		if !isDbConnectionAlive {
			connectToDb(config)
		}
	})
}

func connectToDb(config *appTypes.DatabaseConfig) {

	connectionURI := mongodbConnString(config)
	log.Println(connectionURI)

	ctx := helpers.GetDbConnectContext()

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create MongoDb client: %v", err)
		panic(err)
	}

	err = client.Connect(ctx.Ctx)
	if err != nil {
		log.Printf("Failed to connect to MongoDb cluster: %v", err)
		panic(err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx.Ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
		panic(err)
	} else {
		log.Println("MongoDb: Ping-pong")
	}

	_db = client.Database(config.DbName)
	isDbConnectionAlive = true

	// collection := _db.Collection("issues")
	// //Perform InsertOne operation & validate against the error.
	// _, err = collection.InsertOne(context.TODO(), bson.M{"name": "pi", "value": 3.14159})

	log.Printf("Connected to MongoDB!")

	log.Println("injecting db instance to repository")
	// db instance injection
	repository.UseDb(_db)

	ctx.Cancel()
}

func mongodbConnString(config *appTypes.DatabaseConfig) string {
	// for localhost: mongodb://<username>:<password>@host/_db-name
	// only localhost will have a port
	// localhost dbs will not include srv
	if config.DbHost == localhost {

		if len(config.DbUsername) == 0 || len(config.DbPassword) == 0 {
			return fmt.Sprintf(mongodbLocal, config.DbHost, config.DbPort)
		}
		return fmt.Sprintf(mongodbLocalAuth, config.DbUsername, config.DbPassword, config.DbHost, config.DbPort)
	}

	// for remote host: mongodb+srv://<username>:<password>@host/_db-name
	// atlas clusters will not provide a port or connect to default port 27017
	return fmt.Sprintf(mongodbSRV, config.DbUsername, config.DbPassword, config.DbHost)
}
