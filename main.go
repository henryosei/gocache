package main

import (
	"flag"
	"fmt"

	cache "github.com/loganrk/go-heap-cache"
)

type CacheInstance struct {
	CacheConfig cache.Cache
}

type Config struct {
	Capacity       int
	Expire         int64
	EvictionPolicy int64
}

func (instance *CacheInstance) CacheGetter(key string) (string, error) {
	return instance.CacheGetter(key)
}

func (instance *CacheInstance) CacheSetter(key, value string) {
	instance.CacheSetter(key, value)
}
func main() {
	var config Config
	cacheInstance := newCacheInstance(&config)

	cacheInstance.Set("key", "value")
	data, _ := cacheInstance.Get("key")
	fmt.Printf("Data from cache: %s\n", data)
}

func newCacheInstance(config *Config) cache.Cache {
	flagParser(config)
	conf := cache.Config{Capacity: config.Capacity, Expire: config.Expire, EvictionPolicy: cache.EVICTION_POLICY_LRU}
	cacheInstance := cache.New(&conf)
	return cacheInstance
}

func flagParser(config *Config) {
	var help string
	flag.StringVar(&help, "h", "", "")
	flag.IntVar(&config.Capacity, "c", 100, "Maximum number of items in the cache")
	flag.Int64Var(&config.Expire, "e", 25, "Expiration time in seconds")
	flag.Int64Var(&config.EvictionPolicy, "p", int64(cache.EVICTION_POLICY_FIFO), "Use cache eviction policy.")
	flag.Parse()
}
