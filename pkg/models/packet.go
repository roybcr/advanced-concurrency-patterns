package models

import (
	"fmt"
	"sync"
)

type Packet struct {
	id      interface{}
	data    []byte
	storage map[string]interface{}
	mu      sync.RWMutex
}

func NewPacket(id interface{}, data []byte) *Packet {
	return &Packet{
		id:   id,
		data: data,
	}
}

func (p *Packet) ID() interface{} {
	return p.id
}

func (p *Packet) Data() []byte {
	return p.data
}

func (p *Packet) Get(key string) (value interface{}, exists bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	value, exists = p.storage[key]
	return
}

func (p *Packet) Set(key string, value interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.storage == nil {
		p.storage = make(map[string]interface{})
	}
	p.storage[key] = value
}

func (p *Packet) MustGet(key string) interface{} {
	if v, ok := p.Get(key); ok {
		return v
	}
	panic(fmt.Errorf("key `%s` does not exist", key))
}

func (p *Packet) Remove(key string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.storage, key)
}
