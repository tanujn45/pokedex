package pokecache

import "time"

func (c *Cache) Add(key string, value []byte) {
	ce := cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = ce
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, key)
		}
	}
}
