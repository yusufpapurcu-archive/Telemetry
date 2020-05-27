package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/yusufpapurcu/Telemetry/database"
	"github.com/yusufpapurcu/Telemetry/socket"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.DisableConsoleColor()
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
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
	url, port := setParams()
	database.Connect(url)

	route := gin.New()
	socket.SetSockets(route)
	if err := route.Run(port); err != nil {
		log.Fatal(err)
	}
}

func setParams() (string, string) {
	var env bool
	var url, port string
	flag.BoolVar(&env, "env", false, "Bool Flag.\nUse for Switch Flags/Envs.\nDefault False(Flag)")
	flag.StringVar(&url, "url", "mongodb://localhost:27017", "String Flag.\nDefine url by Using This.\nDefault mongodb://localhost:27017")
	flag.StringVar(&port, "port", ":8080", "String Flag.\nDefine Port by Using This.\nDefault 8080")
	flag.Parse()
	if env {
		url := os.Getenv("DB_URL")
		port := os.Getenv("TELEMPORT")
		return url, port
	}
	return url, port
}
