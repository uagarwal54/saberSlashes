package main

import "sync"

type (
	KV struct {
		mu   sync.Mutex
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
	return nil
}

func (kv *KV) Get(key []byte) ([]byte, bool) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	val, ok := kv.data[string(key)]
	return val, ok
}
