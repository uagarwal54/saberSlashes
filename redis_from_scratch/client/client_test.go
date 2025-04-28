package client

import (
	"context"
	"testing"
)

func TestSetCommand(t *testing.T) {
	c, err := NewClient("localhost:5000")
	if err != nil {
		t.Fatalf("error while creating client conn: %v\n", err)
	}
	defer c.Close()

	/* The diff between context.Background() and context.TODO() is nothing as both return context.emptyCtx which is an empty struct.
	The only diff that can be seen is that the context.Background() will return context.backgroundCtx struct which inherits context.emptyCtx and hence is an empty struct. context.backgroundCtx implements
	the string interface which returns "context.Background" as string
	whereas context.TODO() will return context.todoCtx struct which inherits context.emptyCtx and hence is an empty struct. context.todoCtx implements the string interface which returns "context.TODO" as string
	*/
	if err := c.Set(context.Background(), "foo", "1"); err != nil {
		t.Fatalf("error: %v\n", err)
	}
}

func TestGetCommand(t *testing.T) {
	c, err := NewClient("localhost:5000")
	if err != nil {
		t.Fatalf("Error while creating client conn: %v\n", err)
	}

	defer c.Close()
	if _, err := c.Get(context.Background(), "foo"); err != nil {
		t.Fatalf("error: %v\n", err)
	}
}

func TestSetAndGetCommand(t *testing.T) {
	c, err := NewClient("localhost:5000")
	if err != nil {
		t.Fatalf("error while creating client conn: %v\n", err)
	}
	defer c.Close()
	ctx := context.Background()
	key, expectedValue := "foo", "1"

	if err := c.Set(ctx, key, expectedValue); err != nil {
		t.Fatalf("error: %v\n", err)
	}

	val, err := c.Get(ctx, key)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	if val != expectedValue {
		t.Fatalf("expected %q, got %q", expectedValue, val)
	}
}
