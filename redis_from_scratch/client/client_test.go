package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("localhost:5000")
	if err != nil {
		log.Fatal("Error while creating client conn: ", err)
	}

	for i := 1; i < 10; i++ {
		/* The diff between context.Background() and context.TODO() is nothing as both return context.emptyCtx which is an empty struct.
		The only diff that can be seen is that the context.Background() will return context.backgroundCtx struct which inherits context.emptyCtx and hence is an empty struct. context.backgroundCtx implements
		the string interface which returns "context.Background" as string
		whereas context.TODO() will return context.todoCtx struct which inherits context.emptyCtx and hence is an empty struct. context.todoCtx implements the string interface which returns "context.TODO" as string
		*/
		fmt.Println("Set =>", fmt.Sprintf("foo_%d", i))
		if err := c.Set(context.Background(), fmt.Sprintf("foo_%d", i), fmt.Sprintf("bar_%d", i)); err != nil {
			log.Fatal("Error: ", err)
		}

		if val, err := c.Get(context.Background(), fmt.Sprintf("foo_%d", i)); err != nil {
			log.Fatal("Error: ", err)
		} else {
			fmt.Println("Value: ", val)
		}
	}
	// fmt.Println(server.kv.data)
	time.Sleep(1 * time.Second)
}
