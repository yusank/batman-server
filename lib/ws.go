package lib

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader - ws upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// allow all connection
		return true
	},
}

// StartWS - start ws
func StartWS(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade err:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	defer ws.Close()

	for {
		mt, b, err := ws.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			return
		}

		msg := new(Message)
		if err = UnmarshalMsg(b, msg); err != nil {
			log.Println("msg parse err:", err, string(b))
			continue
		}

		log.Printf("[recv] %+v \n", msg)
		msg.UnixTime = time.Now().UnixNano() / int64(time.Millisecond)
		b, err = MarshallMsg(msg)
		if err != nil {
			return
		}

		err = ws.WriteMessage(mt, b)
		if err != nil {
			log.Println("write err:", err)
		}
	}
}

func ConnectWS(uri, path string) (c *websocket.Conn, err error) {
	u := url.URL{Scheme: "ws", Host: uri, Path: path}

	c, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	return
}
