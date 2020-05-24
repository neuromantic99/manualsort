package main

import (
	"fmt"
	"net/http"

	"github.com/neuromantic99/manualsort/pkg/mergesort"
	"github.com/neuromantic99/manualsort/pkg/websocket"
)

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {

	// Upgrade http
	//ws, err := websocket.Upgrade(w, r)
  _, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	fmt.Println("Connection Successful")

	// Need this if you want to write from the
	// backend through the websocket
	// Currently causing an issue where only
	// 50 % of button clicks are being read

	//go websocket.Writer(ws)
	//websocket.Reader(ws)
	mergesort.BeginSort()
}

func setupRoutes() {
	// Call the serveWs function when get a http request
	// map our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	setupRoutes()
	fmt.Println("Waiting for connection")
	http.ListenAndServe(":8080", nil)
}

