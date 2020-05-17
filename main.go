package main

import (
	"io"
	"log"
	"os"

	"github.com/yusufpapurcu/Telemetry/socket"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(gin.ReleaseMode)

	route := gin.New()
	route.Use(gin.Recovery())

	socket.SetSockets(route)
	if err := route.Run(); err != nil {
		log.Fatal(err)
	}
}
