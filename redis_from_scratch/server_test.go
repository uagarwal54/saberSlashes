package main

import (
	"context"
	"fmt"
	"log"
	"redis_from_scratch/client"
	"sync"
	"testing"
	"time"
)

func TestServerWithMultipleClients(t *testing.T) {
	var server *Server
	go func() {
		server = NewServer(Config{})
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second * 10)

	nClients := 10
	var wg sync.WaitGroup
	for i := 1; i <= nClients; i++ {
		wg.Add(1)
		go func(it int, wg *sync.WaitGroup) {
			defer wg.Done()
			c, err := client.NewClient("localhost:5000")
			if err != nil {
				log.Fatal("Error while creating client conn: ", err)
			}
			defer c.Close()
			key := fmt.Sprintf("client_foo_%d", it)
			val := fmt.Sprintf("client_bar_%d", it)
			if err := c.Set(context.Background(), key, val); err != nil {
				log.Fatal("Error: ", err)
			}
			if val, err := c.Get(context.Background(), key); err != nil {
				log.Fatal("Error: ", err)
			} else {
				fmt.Printf("Client %d got this value back: %s", it, val)
			}
		}(i, &wg)
	}
	wg.Wait()
	log.Println(len(server.peers))
	if len(server.peers) != 0 {
		log.Fatal("Expected 0 peers but got: ", len(server.peers))
	}
}
