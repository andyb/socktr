package main

import (
	"net/http"

	"github.com/apex/log"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", root)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Error: %v", err)
	}

	defer conn.Close()

	for {
		mt, data, err := conn.ReadMessage()
		if err != nil {
			log.Errorf("Error: %v", err)
		}

		if err = conn.WriteMessage(mt, data); err != nil {
			log.Errorf("Error: %v", err)
		}
	}
}
