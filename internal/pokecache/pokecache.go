package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mut   sync.Mutex
	//interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{}
	c.cache = make(map[string]cacheEntry)
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.reapLoop(interval)
		}
	}()
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	b, ok := c.cache[key]
	return b.val, ok
}

func (c *Cache) Add(key string, val []byte) {
	ce := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = ce
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mut.Lock()
	c.mut.Unlock()
	for key, ce := range c.cache {
		if time.Since(ce.createdAt) > interval {

			delete(c.cache, key)

		}
	}
}
