package pokecache

import (
	"sync"
	"time"
)

// create a Map of cacheEntry  and use a mutexe on some functions to prevent them from being modified by many routines.

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheEntries: make(map[string]cacheEntry),
		mu:           &sync.Mutex{},
	}
	cache.ReapLoop(interval)

	return cache
}

func (cache Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (cache Cache) Get(key string) ([]byte, bool) {
	result, ok := cache.cacheEntries[key]
	return result.val, ok
}

func (cache Cache) ReapLoop(interval time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	ticker := time.NewTicker(interval)
	currentTime := <-ticker.C
	for key, entry := range cache.cacheEntries {
		if entry.createdAt.Before(currentTime) {
			delete(cache.cacheEntries, key)
		}
	}
}
