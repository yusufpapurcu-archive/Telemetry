package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var col *mongo.Collection // Paket içinden collection'a ulaşmak için bir nesne oluşturdum.
var logger *log.Logger

//Connect fonksiyonu server'a bağlanıp bize collection verecek.
func Connect(dburl string) {
	// Server'a bağlanma kısmı. RetryWrites kısmı MLab bağlantısı için kapalı. Açmanız sorun olmaz.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburl).SetRetryWrites(false))
	if err != nil {
		log.Fatal("error : " + err.Error())
	}

	// Server bağlantısını test ediyorum.
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error while pinging to the DB", err)
	}

	// İstediğim collectionu alıyorum.
	collection := client.Database("login").Collection("data")
	col = collection
}

// SetLoggerDB function will be get logger structs from main
func SetLoggerDB(log *log.Logger) {
	logger = log
}
