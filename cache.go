package cache

import "time"

type CacheItem struct {
	value    string
	deadline *time.Time
}

func CreateCacheItem(value string, deadline *time.Time) CacheItem {
	return CacheItem{
		value:    value,
		deadline: deadline,
	}
}

type Cache struct {
	storage map[string]CacheItem
}

func NewCache() Cache {
	return Cache{
		storage: map[string]CacheItem{},
	}
}

func (c *Cache) Get(key string) (string, bool) {
	currentTime := time.Now()
	val, ok := c.storage[key]
	if ok && (val.deadline == nil || currentTime.After(*val.deadline)) {
		return val.value, true
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	item := CreateCacheItem(value, nil)
	c.storage[key] = item
}

func (c *Cache) Keys() []string {
	currentTime := time.Now()
	result := make([]string, 0)
	for k, v := range c.storage {
		if v.deadline == nil || currentTime.After(*v.deadline) {
			result = append(result, k)
		}
	}
	return result
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	item := CreateCacheItem(value, &deadline)
	c.storage[key] = item
}
