package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var col *mongo.Collection // Create Global Variable for Share data collection in this package
var logger *log.Logger

//Connect function for Connect database and get collections
func Connect(dburl string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburl).SetRetryWrites(false)) // Connect Database
	if err != nil {
		log.Fatal("error : " + err.Error())
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Error while pinging to the DB", err)
	}
	collection := client.Database("login").Collection("data") //Getting data collection from database
	col = collection                                          // Send collection to global data variable.
}

// SetLoggerDB function will be get logger structs from main
func SetLoggerDB(log *log.Logger) {
	logger = log
}
