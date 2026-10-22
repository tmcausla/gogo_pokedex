package pokecache
import (
	"time"
	"sync"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu 		 sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val 	  []byte
}

func NewCache(newInterval time.Duration) Cache {
	newCache := Cache{
		entries: map[string]cacheEntry{},
		mu: sync.Mutex{},
		interval: newInterval,
	}
	go newCache.reapLoop(newInterval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if (time.Since(entry.createdAt) > duration) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}