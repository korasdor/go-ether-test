package cache

import (
	"sync"
	"time"
)

type item struct {
	value     interface{}
	createdAt int64
	ttl       int64
}

type MemoryCache struct {
	sync.RWMutex
	cache map[interface{}]*item
}

func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{
		cache: make(map[interface{}]*item),
	}

	go c.setTtlTimer()

	return c
}

func (c *MemoryCache) setTtlTimer() {
	for {
		c.Lock()
		for k, v := range c.cache {
			if time.Now().Unix()-v.createdAt > v.ttl {
				delete(c.cache, k)
			}
		}
		c.Unlock()

		<-time.After(time.Second)
	}
}

func (c *MemoryCache) Set(key, value interface{}, ttl int64) error {
	c.Lock()
	c.cache[key] = &item{
		value:     value,
		createdAt: time.Now().Unix(),
		ttl:       ttl,
	}
	c.Unlock()

	return nil
}

func (c *MemoryCache) Get(key interface{}) (interface{}, error) {
	c.RLock()
	item, ok := c.cache[key]
	c.RUnlock()

	if !ok {
		return nil, ErrItemNotFound
	}

	return item.value, nil
}
