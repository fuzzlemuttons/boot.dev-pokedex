package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	interval time.Duration
	mu       sync.Mutex
	entries  map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		mu:      sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, exists := c.entries[key]
	if exists {
		return e.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for k, v := range c.entries {
				if v.createdAt.Add(interval).After(now) {
					delete(c.entries, k)
				}
			}
			c.mu.Unlock()
		}
	}
}
