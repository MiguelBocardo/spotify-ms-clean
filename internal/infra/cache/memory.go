package cache

import (
	"sync"
	"time"
)

type entry struct {
	val interface{}
	exp time.Time
}

type InMemoryCache struct {
	mu   sync.RWMutex
	data map[string]entry
	ttl  time.Duration
}

func NewInMemoryCache(ttl time.Duration) *InMemoryCache {
	return &InMemoryCache{data: map[string]entry{}, ttl: ttl}
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock(); defer c.mu.RUnlock()
	if e, ok := c.data[key]; ok {
		if time.Now().Before(e.exp) {
			return e.val, true
		}
	}
	return nil, false
}

func (c *InMemoryCache) Set(key string, val interface{}) {
	c.mu.Lock(); defer c.mu.Unlock()
	c.data[key] = entry{val: val, exp: time.Now().Add(c.ttl)}
}
