package main

import (
	"fmt"
	"log"
	"net/http"

	websocket "github.com/gorilla/websocket"
)

var _ = websocket.BinaryMessage
var _ = log.Fatal

func setupRoutes() {
	http.HandleFunc(
		"/",
		func(rw http.ResponseWriter, r *http.Request) {
			fmt.Fprint(rw, "Simple Server")
		},
	)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServeTLS(":8080", "", "", nil)
}
