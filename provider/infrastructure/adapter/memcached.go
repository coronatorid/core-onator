package adapter

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/coronatorid/core-onator/provider"
)

//go:generate mockgen -source=./memcached.go -destination=./mocks/memcached_mock.go -package mockAdapter

// Memcached wrap memcache into cache interface
type Memcached struct {
	client MemcachedClient
}

// MemcachedClient wrap default memcache client
type MemcachedClient interface {
	Set(item *memcache.Item) error
	Get(key string) (item *memcache.Item, err error)
}

// AdaptMemcache adapt Cache interface
func AdaptMemcache(client MemcachedClient) *Memcached {
	return &Memcached{client: client}
}

// Set cache value
func (m *Memcached) Set(ctx provider.Context, key string, value []byte, expiration time.Duration) error {
	return m.client.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(int(expiration.Seconds())),
	})
}

// Get cache value
func (m *Memcached) Get(ctx provider.Context, key string) (provider.CacheItem, error) {
	i, err := m.client.Get(key)
	if err == memcache.ErrCacheMiss {
		return nil, provider.ErrCacheMiss
	} else if err != nil {
		return nil, err
	}

	return NewCacheItem(i.Key, i.Value, time.Duration(i.Expiration)*time.Second), nil
}
