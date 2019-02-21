package lib

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader - ws upgrader
var upgrader = websocket.Upgrader{}

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
		}

		var msg *Message
		if err = UnmarshalMsg(b, msg); err != nil {
			log.Println("smg parse err:", err)
			continue
		}

		log.Printf("[recv] %+v \n", msg)

		err = ws.WriteMessage(mt, b)
		if err != nil {
			log.Println("write err:", err)
		}
	}
}
