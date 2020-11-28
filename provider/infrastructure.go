package provider

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/coronatorid/core-onator/entity"
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

// Network are external connecion for infrastructure
type Network interface {
	// successBinder must json assignable struct
	GET(ctx context.Context, cfg NetworkConfig, path string, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError
	POST(ctx context.Context, cfg NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError
}

// NetworkConfig given for network request
type NetworkConfig interface {
	Host() string
	Username() string
	Password() string
	Timeout() time.Duration
	Retry() int
	RetrySleepDuration() time.Duration
}
