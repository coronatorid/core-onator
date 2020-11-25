package provider

import (
	"context"
	"errors"
	"time"
)

//go:generate mockgen -source=./infrastructure.go -destination=./mocks/infrastructure_mock.go -package mockProvider

// ErrCacheMiss returned when value from cache is not found
var ErrCacheMiss = errors.New("cache miss")

// Cache is interface to connect to cache infrastructure
type Cache interface {
	// To make it not expire set expiration into 0
	Set(ctx context.Context, key string, value []byte, expiration time.Duration) error
	Get(ctx context.Context, key string) (CacheItem, error)
}

// CacheItem contain result get from Cache interface
type CacheItem interface {
	Key() string
	Value() []byte
	ExpiresIn() time.Duration
}

// TextPublisher handle whatsapp message and maybe sms in the future
type TextPublisher interface {
	Publish(ctx context.Context, phoneNumber, message string) error
}
