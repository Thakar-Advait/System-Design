package main

import (
	"fmt"

	"example.com/singletonmethod/cache"
)

func main() {
	cacheIns1 := cache.GetCacheInstance()
	cacheIns2 := cache.GetCacheInstance()
	fmt.Println(cacheIns2.Store("key1", "value1"))
	fmt.Printf("Same instance: %v\n", cacheIns1 == cacheIns2)
}
