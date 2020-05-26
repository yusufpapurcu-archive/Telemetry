package database

import (
	"context"
	"time"

	"github.com/yusufpapurcu/Telemetry/models"
)

//WriteDataFrame function will save our data to MongoDB server
func WriteDataFrame(data models.SolidData) {
	data.CreatedAt = time.Now().Unix()
	res, err := col.InsertOne(context.TODO(), data) // Create Function
	if err != nil {
		logger.Println(err)
		return
	}
	if res == nil {
		logger.Println("bos")
		return
	}
	logger.Println(res.InsertedID)
}
