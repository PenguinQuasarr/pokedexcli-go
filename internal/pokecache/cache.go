package pokecache

import (
	"sync"
	"time"
)


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
	
}

func NewCache(interval time.Duration) Cache {
	return Cache {
		cache: make(map[string]cacheEntry),
		mu: sync.Mutex{},
		interval: interval,
	}
}

func (ca *Cache) Add(key string, val []byte) {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	ca.cache[key] = cacheEntry {
		createdAt: time.Now(),
		val: val,
	}
}

func (ca *Cache) Get(key string) ([]byte, bool) {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	val, ok := ca.cache[key]
	if !ok {
		return []byte{}, false
	}
	return val.val, true
}

func (ca *Cache) reapLoop() {
	ticker := time.NewTicker(ca.interval)

	ca.mu.Lock()
	defer ca.mu.Unlock()



}
