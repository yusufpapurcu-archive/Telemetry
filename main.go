package main

import (
	"io"
	"log"
	"os"

	"github.com/yusufpapurcu/Telemetry/database"
	"github.com/yusufpapurcu/Telemetry/socket"

	"github.com/gin-gonic/gin"
)

func init() {
	f, err := os.Create("logs/log.txt")
	if err != nil {
		log.Println(err)
	}
	logger := log.New(f, "", log.Ltime)
	logger.SetFlags(log.Lshortfile | log.Ltime)
	database.SetLoggerDB(logger)

	f, err = os.Create("logs/data.txt")
	if err != nil {
		logger.Println("Create Data Log Error: ", err)
	}
	dataSaver := log.New(f, "", log.Lmsgprefix)
	dataSaver.SetFlags(log.Ldate | log.Ltime)
	socket.SetLoggerSocket(logger, dataSaver)
}

func main() {
	gin.DisableConsoleColor()
	database.Connect("mongodb://Test:test12@ds141674.mlab.com:41674/login")
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	route := gin.New()
	route.Use(gin.Recovery())

	socket.SetSockets(route)
	if err := route.Run(); err != nil {
		log.Fatal(err)
	}
}
