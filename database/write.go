package database

import (
	"context"
	"log"
	"time"

	"github.com/yusufpapurcu/Telemetry/models"
)

func WriteDataFrame(data models.SolidData) {
	data.CreatedAt = time.Now().Unix()
	res, err := col.InsertOne(context.TODO(), data) // Create Function
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
