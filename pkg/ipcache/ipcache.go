package ipcache

import (
	"log/slog"
	"s-coder-snippet-sharder/pkg/background"
	"sync"
	"time"
)

const itemLifeMinutes = 1

type IPCacheItem struct {
	InitialRequestTime time.Time
	LastSeenAt         time.Time
}

type Cache struct {
	cache map[string]*IPCacheItem
	mu    sync.RWMutex
}

func New() *Cache {
	c := &Cache{
		cache: make(map[string]*IPCacheItem),
		mu:    sync.RWMutex{},
	}

	//Clear the ip cache every 5 minute
	background.Go(backgroundCacheClearer(c))

	return c
}

func backgroundCacheClearer(c *Cache) func() {
	return func() {
		for range time.Tick(30 * time.Second) {
			func() {
				c.mu.Lock()
				defer c.mu.Unlock()

				slog.Info("Checking IP cache for stale IP's")
				//Every 30 seconds, check every ip in the cache to see if it can be deleted.
				for ip, item := range c.cache {
					if item.InitialRequestTime.Before(time.Now().Add(-itemLifeMinutes * time.Minute)) {
						slog.Info("found ip to be deleted", "ip", ip)
						delete(c.cache, ip)
					}
				}
			}()
		}
	}
}

// Set adds the item to the cache
func (c *Cache) Set(ip string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[ip] = &IPCacheItem{
		InitialRequestTime: time.Now(),
	}
}

// Has return true if the ip exists in the cache
func (c *Cache) Has(ip string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, ok := c.cache[ip]
	if ok {
		item.LastSeenAt = time.Now()
	}
	return ok
}
