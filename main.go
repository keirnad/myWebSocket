package websocket

import (
	"github.com/gorilla/websocket"
)

// Create upgrader to update http connections to websocket protocol
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
