package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]cacheEntry
	mu   sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	new_cache := &Cache{
		data: make(map[string]cacheEntry),
	}
	return new_cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	result,ok := c.data[key]
	if !ok {
		return nil, false
	}
	return result.val, true
}

