package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	log.Printf("Connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	// Reader
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// Send a message
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!"))
	if err != nil {
		log.Println("Write error:", err)
		return
	}

	// Wait or interrupt
	select {
	case <-done:
	case <-interrupt:
		log.Println("Interrupt received, closing connection...")
		_ = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bye"))
		time.Sleep(time.Second)
	}
}
