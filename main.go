package main

import (
	"io"
	"log"
	"os"

	"github.com/yusufpapurcu/Telemetry/database"
	"github.com/yusufpapurcu/Telemetry/socket"

	"github.com/gin-gonic/gin"
)

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
