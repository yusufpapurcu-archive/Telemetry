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

// Listenerları tutmak için waiters kanalını açtım.
var waiters = make(map[*websocket.Conn]bool)

// Veri akışını sağlamak için broadcast kanalını açtım.
var broadcast = make(chan []byte)

// Logger objelerini çekmek için burada da tanımladım.
var logger, dataSaver *log.Logger

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

	go manager()
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
			break
		}
		// Unmarshal işlemi için obje oluşturdum.
		var data models.SolidData
		err = json.Unmarshal(message, &data)
		if err != nil {
			logger.Println(err)
			continue
		}

		// Burada da aldığım veriyi hem dosyaya hem de MongoDB serverine yolladım.
		dataSaver.Println(string(message))
		database.WriteDataFrame(data)

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

// SetLoggerSocket loggerları ana paketten çeken fonksiyon. Daha iyisini bulunca güncelleyeceğim.
func SetLoggerSocket(log, ds *log.Logger) {
	logger, dataSaver = log, ds
}
