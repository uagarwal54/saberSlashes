package main

import (
	"log"
	"sync"
)

type (
	KV struct {
		mu   sync.RWMutex
		data map[string][]byte
	}
)

func NewKV() KV {
	return KV{
		data: map[string][]byte{},
	}
}

func (kv *KV) Set(key, val []byte) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[string(key)] = val
	log.Printf("|| Set || key: %s;val: %s ||\n", string(key), string(val))
	return nil
}

func (kv *KV) Get(key []byte) ([]byte, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	val, ok := kv.data[string(key)]
	log.Printf("|| Get || key: %s;val: %s ||\n", string(key), string(val))
	return val, ok
}
