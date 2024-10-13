package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"redis_from_scratch/client"
	"sync"
	"testing"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func TestServerWithSingleClient(t *testing.T) {
	// Setup the server and wait for it to initialize
	server := NewServer(Config{})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second)

	ctx := context.Background()
	setAndGetFunc := func() error {
		c, err := client.NewClient("localhost:5001")
		if err != nil {
			return err
		}
		defer c.Close()

		key, expectedValue := "foo", "bar"
		if err := c.Set(ctx, key, expectedValue); err != nil {
			return err
		}

		actualValue, err := c.Get(ctx, key)
		if err != nil {
			return err
		}

		if actualValue != expectedValue {
			return fmt.Errorf("expected value %q, got %q", expectedValue, actualValue)
		}

		slog.Info("[TestServerWithSingleClient]", slog.String("key", key), slog.String("value", actualValue))
		return nil
	}
	if err := setAndGetFunc(); err != nil {
		t.Errorf("error in setting and getting from cache: %v\n", err)
		return
	}

	// To release the connected peers
	time.Sleep(time.Second)
	if len(server.peers) != 0 {
		t.Errorf("expected 0 peers, got: %d\n", len(server.peers))
		return
	}
}

func TestServerWithMultipleClients(t *testing.T) {
	// Setup the server and wait for it to initialize
	server := NewServer(Config{})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second)

	// Store the key-pair with client info
	type multipleSetAndGet struct {
		client     int
		key, value string
	}
	nClients := 10
	multipleSetAndGetPairs := make([]multipleSetAndGet, nClients)

	var wg sync.WaitGroup
	for i := 1; i <= nClients; i++ {
		wg.Add(1)
		go func(it int, wg *sync.WaitGroup) {
			defer wg.Done()
			c, err := client.NewClient("localhost:5001")
			if err != nil {
				t.Errorf("error while creating client conn: %v\n", err)
				return
			}
			defer c.Close()
			ctx := context.Background()

			key := fmt.Sprintf("client_foo_%d", it)
			expectedvalue := fmt.Sprintf("client_bar_%d", it)
			if err := c.Set(ctx, key, expectedvalue); err != nil {
				t.Errorf("error while setting the key in cache: %v\n", err)
				return
			}

			actualValue, err := c.Get(ctx, key)
			if err != nil {
				t.Errorf("error while retrieving value from cache: %v\n", err)
				return
			}

			if actualValue != expectedvalue {
				t.Errorf("client %d expected %q, got %q", it, expectedvalue, actualValue)
				return
			}

			multipleSetAndGetPairs[it-1] = multipleSetAndGet{client: it, key: key, value: actualValue}
		}(i, &wg)
	}
	wg.Wait()

	// To release the connected peers
	time.Sleep(time.Second)
	if len(server.peers) != 0 {
		t.Errorf("expected 0 peers, got: %d\n", len(server.peers))
		return
	}

	for _, setAndGetPair := range multipleSetAndGetPairs {
		slog.Info("[TestServerWithMultipleClients]",
			slog.Int("client", setAndGetPair.client),
			slog.String("key", setAndGetPair.key),
			slog.String("value", setAndGetPair.value),
		)
	}
}

// This won't work, as we're not following redis protocol.
// We'll need to use their response handling protocol to make this work.
// We're implementing basic redis functionalities for now.
func TestServerWithOfficialRedisClient(t *testing.T) {
	server := NewServer(Config{})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Second)
	// Create a context for Redis operations
	ctx := context.Background()

	// Connect to the Redis server
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:5001", // Redis server address
		Password: "",               // No password set (default)
		DB:       0,                // Use default DB
	})
	defer rdb.Close()

	key, value := "exampleKey", "Hello, Redis!"
	// Set a key-value pair in Redis
	if err := rdb.Set(ctx, key, value, 0).Err(); err != nil {
		t.Errorf("failed to set key: %v\n", err)
		return
	}

	// Get the value for the key from Redis
	fetchedVal, err := rdb.Get(ctx, key).Result()
	if err != nil {
		t.Errorf("failed to get key: %v\n", err)
		return
	}

	if fetchedVal != value {
		t.Errorf("expected %s, got %s", value, fetchedVal)
		return
	}
}

func TestFooBar(t *testing.T) {
	in := map[string]string{
		"server":  "redis",
		"version": "6.0",
	}
	out := respWriteMap(in)
	fmt.Println(string(out))
}
