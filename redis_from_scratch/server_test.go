package main

import (
	"context"
	"fmt"
	"log"
	"redis_from_scratch/client"
	"sync"
	"testing"
	"time"

	redis "github.com/redis/go-redis/v9"
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

func TestServerWithOfficialRedisClient(t *testing.T) {
	var server *Server
	go func() {
		server = NewServer(Config{})
		log.Fatal(server.Start())
	}()
	fmt.Println("Waiting for the server to start...")
	time.Sleep(time.Second * 10)
	// Create a context for Redis operations
	var ctx = context.Background()

	// Connect to the Redis server
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:5001", // Redis server address
		Password: "",               // No password set (default)
		DB:       0,                // Use default DB
	})
	fmt.Println(rdb)
	fmt.Println("This is working")
	key := "exampleKey"
	val := "Hello, Redis!"
	// Set a key-value pair in Redis
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set key: %v", err)
	}

	// Get the value for the key from Redis
	fetchedVal, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}
	if fetchedVal != val {
		t.Fatalf("Expected %s but got %s", val, fetchedVal)
	}
	fmt.Println("exampleKey:", val)

}

func TestFooBar(t *testing.T) {
	in := map[string]string{
		"server":  "redis",
		"version": "6.0",
	}
	out := respWriteMap(in)
	fmt.Println(string(out))
}
