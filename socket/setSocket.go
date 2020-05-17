package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var waiters = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SetSockets a main function of this package
func SetSockets(route *gin.Engine) {

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
				log.Fatal(err)
				waiters[ws] = false
			}
		}
	}
}

// ListenerForCar function for manage cars
func ListenerForCar(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	var data solidData
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = json.Unmarshal(message, &data)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		fmt.Println(data)
		broadcast <- message
	}

}

// DataWaiters for manage data-wanter platform's
func DataWaiters(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	// register waiter
	waiters[ws] = true
}
