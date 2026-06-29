# Simple Go WebSocket Server

A minimal WebSocket echo server built in Go using the `gorilla/websocket` package. It upgrades HTTP connections to WebSockets, reads incoming messages, logs them, and echoes them back to the client.

## Setup & Installation

1. **Initialize the module** (if not done yet):
   ```bash
   go mod init websocket-server
   ```

2. **Install dependencies:**
   ```bash
   go get github.com/gorilla/websocket
   ```

3. **Run the server:**
   ```bash
   go run main.go
   ```

## How to Test

Open your browser's Developer Tools (`F12`), go to the **Console** tab, and run:

```javascript
const ws = new WebSocket("ws://localhost:8080/ws");

ws.onopen = () => ws.send("Hello Server!");
ws.onmessage = (event) => console.log("Echo from server:", event.data);
```