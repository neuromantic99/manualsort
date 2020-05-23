package main

import (
	"fmt"
	"net/http"

	"github.com/neuromantic99/manualsort/pkg/websocket"
)

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {

	// Upgrade http
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	// Need this if you want to write from the
  // backend through the websocket
	// Currently causing an issue where only
	// 50 % of button clicks are being read 
	//go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	// Call the serveWs function when get a http request
	// map our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Howdy from manual-sort")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
