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
	store map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		store: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.store[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.store[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		cache.mu.Lock()
		for key, entry := range cache.store {
			if interval < time.Since(entry.createdAt) {
				delete(cache.store, key)
			}
		}
		cache.mu.Unlock()
	}
}
