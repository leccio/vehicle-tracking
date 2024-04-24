package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
	"vehicle-tracking/backend/types"

	"github.com/gorilla/websocket"
)

func init() {
	rand.New(rand.NewSource(time.Now().Unix()))
}

// Simulates n(vehicles) producing data
func main() {
	vehicles := 3

	for range vehicles {
		obu := newObu()

		conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/ws", nil)
		if err != nil {
			log.Fatal(err)
		}

		go func(conn *websocket.Conn) {
			for {
				lat, long := generateCoords()
				obu.Lat = lat
				obu.Long = long

				fmt.Printf("%+v\n", obu)
				if err := conn.WriteJSON(obu); err != nil {
					log.Fatal(err)
				}

				time.Sleep(time.Millisecond * 400)
			}
		}(conn)
	}

	for {
	}
}

func newObu() *types.Obu {
	return &types.Obu{Id: rand.Intn(math.MaxInt)}
}

func generateCoords() (float64, float64) {
	rlat := float64(rand.Intn(100) + 1)
	rlong := float64(rand.Intn(100) + 1)
	flat := rand.Float64()
	flong := rand.Float64()
	return rlat + flat, rlong + flong
}
