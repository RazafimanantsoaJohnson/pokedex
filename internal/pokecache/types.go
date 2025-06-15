package pokecache

import (
	"sync"
	"time"
)

// Will hold our list of 'cacheEntries' and protect them from being mutated from many places(threads)
type Cache struct {
	mu           *sync.Mutex
	cacheEntries map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
