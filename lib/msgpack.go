package lib

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader - ws upgrader
var upgrader = websocket.Upgrader{}

// Start - start ws
func Start(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer ws.Close()

	for {
		mt, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
		}

		log.Println("got", msg)
		err = ws.WriteMessage(mt, msg)
		if err != nil {
			log.Println("write err:", err)
		}
	}
}
