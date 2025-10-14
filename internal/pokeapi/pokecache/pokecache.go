package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mutex sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entry: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	// lock so no one can touch while we add
	c.mutex.Lock()
	// unlock after we leave the function
	defer c.mutex.Unlock()
	//add new data to the cache
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	// lock before getting
	c.mutex.Lock()
	defer c.mutex.Unlock()
	e, ok := c.entry[key]
	if !ok {
		return nil, false
	}
	return e.val, true
}

func (c *Cache) reapLoop() {
	// make a timer
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.mutex.Lock()
		for k, e := range c.entry {
			// is expired one inerval after it was inserted
			expiresAt := e.createdAt.Add(c.interval)
			// the entry is older than the interval if true, so delete
			if expiresAt.Before(now) {
				delete(c.entry, k)
			}
		}
		c.mutex.Unlock()

	}
}