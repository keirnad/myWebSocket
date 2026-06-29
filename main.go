package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Create upgrader to update http connections to websocket protocol
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Recieved: %s", message)

		if err := ws.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			break
		}
	}
}
