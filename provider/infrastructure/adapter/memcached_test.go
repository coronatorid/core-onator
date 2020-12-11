package adapter_test

import (
	"errors"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/infrastructure/adapter"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/stretchr/testify/assert"

	mockAdapter "github.com/coronatorid/core-onator/provider/infrastructure/adapter/mocks"

	"github.com/golang/mock/gomock"
)

func TestMemcached(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("Set", func(t *testing.T) {
		memcacheClient := mockAdapter.NewMockMemcachedClient(mockCtrl)
		memcacheClient.EXPECT().Set(gomock.Any()).Return(nil)

		memcachedAdapter := adapter.AdaptMemcache(memcacheClient)

		assert.Nil(t, memcachedAdapter.Set(testhelper.NewTestContext(), "key", []byte("value"), 0))
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("It return CacheItem", func(t *testing.T) {
			key := "key"
			value := []byte("value")
			expiration := time.Second

			memcacheClient := mockAdapter.NewMockMemcachedClient(mockCtrl)
			memcacheClient.EXPECT().Get(key).Return(&memcache.Item{
				Key:        key,
				Value:      value,
				Expiration: int32(int(expiration.Seconds())),
			}, nil)

			memcachedAdapter := adapter.AdaptMemcache(memcacheClient)

			cacheItem, err := memcachedAdapter.Get(testhelper.NewTestContext(), key)
			assert.Nil(t, err)
			assert.Equal(t, key, cacheItem.Key())
			assert.Equal(t, value, cacheItem.Value())
			assert.Equal(t, expiration, cacheItem.ExpiresIn())
		})

		t.Run("When there is unexpected error then it will return the error", func(t *testing.T) {
			memcacheClient := mockAdapter.NewMockMemcachedClient(mockCtrl)
			memcacheClient.EXPECT().Get("key").Return(nil, errors.New("unexpected error"))

			memcachedAdapter := adapter.AdaptMemcache(memcacheClient)

			cacheItem, err := memcachedAdapter.Get(testhelper.NewTestContext(), "key")
			assert.Nil(t, cacheItem)
			assert.NotNil(t, err)
		})

		t.Run("When it's cache not found error then it will return provider.ErrCacheMiss", func(t *testing.T) {
			memcacheClient := mockAdapter.NewMockMemcachedClient(mockCtrl)
			memcacheClient.EXPECT().Get("key").Return(nil, memcache.ErrCacheMiss)

			memcachedAdapter := adapter.AdaptMemcache(memcacheClient)

			cacheItem, err := memcachedAdapter.Get(testhelper.NewTestContext(), "key")
			assert.Nil(t, cacheItem)
			assert.Equal(t, provider.ErrCacheMiss, err)
		})
	})
}
