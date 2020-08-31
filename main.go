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
	// Log çıkarırken renklere ihtiyaç yok.
	gin.DisableConsoleColor()

	// Default logger için dosya oluşturup gin'e verdim.
	f, _ := os.OpenFile("logs/gin.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// Hata ve bilgilerin yazılması için log dosyası açtım.
	f, err := os.OpenFile("logs/log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
	}
	// Dosyayı kullanarak yeni bir logger objesi oluşturdum ve çıktı biçimini düzenledim.
	logger := log.New(f, "", log.Ltime)
	logger.SetFlags(log.Lshortfile | log.Ltime)

	// Database'e logger'ını verdim.
	database.SetLoggerDB(logger)

	// Veri kayıtlarını ilkel biçimde de tutuyorum.
	f, err = os.OpenFile("logs/data.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		logger.Println("Create Data Log Error: ", err)
	}

	// Dosyayı kullanarak yeni bir logger objesi oluşturdum ve çıktı biçimini düzenledim.
	dataSaver := log.New(f, "", log.Lmsgprefix)
	dataSaver.SetFlags(log.Ldate | log.Ltime)
	// Kayıt için logger'ını verdim.
	socket.SetLoggerSocket(logger, dataSaver)
}

func main() {
	_, port := setParams()
	//database.Connect(url)

	route := gin.New()
	socket.SetSockets(route)
	if err := route.Run(port); err != nil {
		log.Fatal(err)
	}
}

// setParams fonksiyonu çalıştırırken temel bilgileri almak için.
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
