package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yusufpapurcu/Telemetry/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yusufpapurcu/Telemetry/database"
)

var waiters = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)
var logger, dataSaver *log.Logger
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SetSockets a main function of this package
func SetSockets(route *gin.Engine) {
	route.Use(gin.Recovery())

	route.GET("/data/listen", DataWaiters)
	route.GET("/data/post", ListenerForCar)

	go manager()
}

func manager() {
	for {
		a := <-broadcast
		for ws := range waiters {
			if !waiters[ws] {
				ws.Close()
				delete(waiters, ws)
				continue
			}
			if err := ws.WriteMessage(1, a); err != nil {
				logger.Println("Brodcast Error :", err)
				waiters[ws] = false
			}
		}
	}
}

// ListenerForCar function for manage cars
func ListenerForCar(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Println(err)
	}
	defer ws.Close()
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			logger.Println("Read Error:", err)
			break
		}
		var data models.SolidData
		err = json.Unmarshal(message, &data)
		if err != nil {
			logger.Println(err)
			continue
		}
		dataSaver.Println(string(message))
		database.WriteDataFrame(data)
		broadcast <- message
	}

}

// DataWaiters for manage data-wanter platform's
func DataWaiters(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Println("Logging Error", err)
	}

	// register waiter
	waiters[ws] = true
}

// SetLoggerSocket function will be get logger structs from main
func SetLoggerSocket(log, ds *log.Logger) {
	logger, dataSaver = log, ds
}
