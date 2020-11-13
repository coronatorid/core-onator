package provider

import (
	"errors"
	"time"
)

//go:generate mockgen -source=./infrastructure.go -destination=./mocks/infrastructure_mock.go -package mockProvider

// ErrCacheMiss returned when value from cache is not found
var ErrCacheMiss = errors.New("cache miss")

// Cache is interface to connect to cache infrastructure
type Cache interface {
	// To make it not expire set expiration into 0
	Set(key string, value []byte, expiration time.Duration) error
	Get(key string) (CacheItem, error)
}

// CacheItem contain result get from Cache interface
type CacheItem interface {
	Key() string
	Value() []byte
	ExpiresIn() time.Duration
}
