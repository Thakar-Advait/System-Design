package cache

import (
	"fmt"
	"sync"
)

type Cache struct{}

var cache *Cache
var once sync.Once

func GetCacheInstance() *Cache {
	once.Do(func() {
		fmt.Println("Creating a cache instance...")
		cache = &Cache{}
	})
	return cache
}

func (*Cache) Store(key, val string) string {
	return fmt.Sprintf("Storing in cache: %s:%s", key, val)
}
