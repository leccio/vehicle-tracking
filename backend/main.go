package main

import (
	"log"
	"net/http"
	"vehicle-tracking/backend/types"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println("client connected")
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1028,
			WriteBufferSize: 1028,
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		for {
			var data types.Obu
			if err := conn.ReadJSON(&data); err != nil {
				log.Fatal(err)
			}
			log.Println(data)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
