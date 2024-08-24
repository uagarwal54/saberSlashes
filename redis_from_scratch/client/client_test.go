package client

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestSetCommand(t *testing.T) {
	c, err := NewClient("localhost:5000")
	if err != nil {
		log.Fatal("Error while creating client conn: ", err)
	}
	defer c.Close()
	/* The diff between context.Background() and context.TODO() is nothing as both return context.emptyCtx which is an empty struct.
	The only diff that can be seen is that the context.Background() will return context.backgroundCtx struct which inherits context.emptyCtx and hence is an empty struct. context.backgroundCtx implements
	the string interface which returns "context.Background" as string
	whereas context.TODO() will return context.todoCtx struct which inherits context.emptyCtx and hence is an empty struct. context.todoCtx implements the string interface which returns "context.TODO" as string
	*/
	if err := c.Set(context.Background(), "foo", "1"); err != nil {
		log.Fatal("Error: ", err)
	}
}

func TestGetCommand(t *testing.T) {
	c, err := NewClient("localhost:5000")
	if err != nil {
		log.Fatal("Error while creating client conn: ", err)
	}
	defer c.Close()
	if val, err := c.Get(context.Background(), "foo"); err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("Value: ", val)
	}

}
