package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mapCacheEntry map[string]cacheEntry
	mu            sync.Mutex
	interval      time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mapCacheEntry: make(map[string]cacheEntry),
		interval:      interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mapCacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.mapCacheEntry[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for timeStamp := range ticker.C {
		c.mu.Lock()

		for key, entry := range c.mapCacheEntry {
			if timeStamp.Sub(entry.createdAt) >= c.interval {
				delete(c.mapCacheEntry, key)
			}
		}

		c.mu.Unlock()
	}
}
