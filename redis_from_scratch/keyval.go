package main

import "sync"

type (
	KV struct {
		mu   sync.RWMutex
		data map[string]string
	}
)

func NewKV() KV {
	return KV{
		data: map[string]string{},
	}
}

func (kv *KV) Set(key, val string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[key] = val
	return nil
}

func (kv *KV) Get(key string) (string, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	val, ok := kv.data[key]
	return val, ok
}
