package main

import (
	"encoding/json"
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

		if err := conn.WriteMessage(websocket.TextMessage, []byte("connected!")); err != nil {
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

	http.HandleFunc("POST /coordinates", func(w http.ResponseWriter, r *http.Request) {
		var data types.Obu
		err := json.NewDecoder(r.Body).Decode(&data)
		defer r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(data)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
