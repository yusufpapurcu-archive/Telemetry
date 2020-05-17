package main

import (
	"log"

	"github.com/yusufpapurcu/Telemetry/socket"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	socket.SetSockets(route)
	if err := route.Run(); err != nil {
		log.Fatal(err)
	}
}
