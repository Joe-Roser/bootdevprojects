package pokecache

import (
	"context"
	"sync"
	"time"
)

type PokeCache struct {
	cache map[string]cacheEntry
	lock  sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       any
}

func NewCache(ctx context.Context, interval time.Duration) *PokeCache {
	cache := PokeCache{
		cache: make(map[string]cacheEntry),
		lock:  sync.Mutex{},
	}

	go cache.reapLoop(ctx, interval)

	return &cache
}

func (c *PokeCache) Add(key string, val any) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *PokeCache) Get(key string) (any, bool) {
	c.lock.Lock()

	val, ok := c.cache[key]

	c.lock.Unlock()

	if ok {
		return val.val, true
	}

	return nil, false
}

func (c *PokeCache) reapLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			c.lock.Lock()
			for key, val := range c.cache {
				if time.Since(val.createdAt) > interval {
					delete(c.cache, key)
				}
			}
			c.lock.Unlock()
		case <-ctx.Done():
			return
		}
	}
}
