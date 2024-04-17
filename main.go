package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

const PORT = 3000

func main() {
	server := newServer()
	http.Handle("/", websocket.Server{
		// Config:  websocket.Config{},

		Handler: websocket.Handler(server.HandleWS),
	})
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

func HandleSockets(w http.ResponseWriter, r *http.Request) {

}
