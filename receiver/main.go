package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"vehicle-tracking/receiver/types"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/websocket"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

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
			// log.Println(data)

			// Produce messages to topic (asynchronously)
			msg, _ := json.Marshal(data)
			topic := "myTopic"
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          msg,
			}, nil)

			// Wait for message deliveries before shutting down
			// p.Flush(15 * 1000)
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

	log.Println("server started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
