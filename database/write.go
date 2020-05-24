package database

import (
	"context"
	"log"
	"time"

	"github.com/yusufpapurcu/Telemetry/models"
)

func WriteDataFrame(data models.SolidData) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Context for Create function
	res, err := col.InsertOne(ctx, data)                               // Create Function
	if err != nil {
		log.Println(err)
		return
	}
	if res == nil {
		log.Print("bos")
		return
	}
	log.Println(res.InsertedID)
}
