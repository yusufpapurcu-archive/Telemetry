package socket

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Listenerları tutmak için waiters kanalını açtım.
var waiters = make(map[*websocket.Conn]bool)
var mobileWaiters = make(map[*websocket.Conn]bool)

// Veri akışını sağlamak için broadcast kanalını açtım.
var broadcast = make(chan []byte)
var mobileBroadcast = make(chan []byte)

// Logger objelerini çekmek için burada da tanımladım.
var logger, dataSaver *log.Logger
var data string

// Websocket için upgrader kısmı. Websocket ayarlarının çoğunu buradan yapabilirsiniz. En basit hali bu.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SetSockets bu paketin ana fonksiyonu
func SetSockets(route *gin.Engine) {
	route.Use(gin.Recovery())

	route.GET("/data/listen", DataWaiters)
	route.GET("/data/post", ListenerForCar)
	route.GET("data/listen/mobile", MobileDataWaiters)
	go manager()
	go mobileManager()
}

// manager dinleyiciler'e data yollayacak olan ve ayrı bir rutinde çalışacak fonksiyon.
func manager() {
	for {
		// Gelen veriyi olduğu gibi dinleyicilere yolluyor bu kısım. Kontrol eklemek isterseniz buraya ekleyebilirsiniz.
		a := <-broadcast
		for ws := range waiters {
			// Soket hata verirse onu listeden siliyor ve soketi kapatıyor.
			if !waiters[ws] {
				ws.Close()
				delete(waiters, ws)
				continue
			}
			if err := ws.WriteMessage(1, a); err != nil {
				// Hata olursa log'u çıkarıyor ve listeyi yeniliyor.
				logger.Println("Brodcast Error :", err)
				waiters[ws] = false
			}
		}
	}
}

// manager dinleyiciler'e data yollayacak olan ve ayrı bir rutinde çalışacak fonksiyon.
func mobileManager() {
	for {
		// Gelen veriyi olduğu gibi dinleyicilere yolluyor bu kısım. Kontrol eklemek isterseniz buraya ekleyebilirsiniz.
		a := <-mobileBroadcast
		for ws := range mobileWaiters {
			// Soket hata verirse onu listeden siliyor ve soketi kapatıyor.
			if !mobileWaiters[ws] {
				ws.Close()
				delete(mobileWaiters, ws)
				continue
			}
			if err := ws.WriteMessage(1, a); err != nil {
				// Hata olursa log'u çıkarıyor ve listeyi yeniliyor.
				logger.Println("Brodcast Error :", err)
				mobileWaiters[ws] = false
			}
		}
	}
}

// ListenerForCar Data yollanacak soket.
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
			ws.Close()
			break
		}

		// Burada da aldığım veriyi hem dosyaya hem de MongoDB serverine yolladım.
		dataSaver.Println(string(message))
		//	database.WriteDataFrame(data)
		if len(message) > 3 {

			if message[3] == byte('3') {
				ws.Close()
				fmt.Println("Closed")
			}
			switch string(message)[0:2] {
			case "11":
				lsit := strings.Split(string(message), "/")
				data += lsit[1] + "*"
				break
			case "12":
				lsit := strings.Split(string(message), "/")
				data += lsit[1] + "*"
				break
			case "26":
				lsit := strings.Split(string(message), "/")
				data += lsit[1] + "*"
				break
			case "23":
				lsit := strings.Split(string(message), "/")
				data += lsit[1]
				mobileBroadcast <- []byte(data)
				data = ""
				break
			}
		}
		// Dinleyicilere gidiyor buradan da.
		broadcast <- message
	}

}

// DataWaiters veri bekleyenleri waiters listesine alan fonksiyon.
func DataWaiters(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Println("Logging Error", err)
	}

	// Kayıt altına alıyorum.
	waiters[ws] = true
}

// DataWaiters veri bekleyenleri waiters listesine alan fonksiyon.
func MobileDataWaiters(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Println("Logging Error", err)
	}

	// Kayıt altına alıyorum.
	mobileWaiters[ws] = true
}

// SetLoggerSocket loggerları ana paketten çeken fonksiyon. Daha iyisini bulunca güncelleyeceğim.
func SetLoggerSocket(log, ds *log.Logger) {
	logger, dataSaver = log, ds
}
