package adapter_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/coronatorid/core-onator/provider/infrastructure/adapter"
	"github.com/dghubble/sling"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type mockNetworkConfig struct {
	host     string
	username string
	password string
	timeout  time.Duration
	retry    int
	sleep    time.Duration
}

func (m *mockNetworkConfig) Host() string {
	return m.host
}

func (m *mockNetworkConfig) Username() string {
	return m.username
}

func (m *mockNetworkConfig) Password() string {
	return m.password
}
func (m *mockNetworkConfig) Timeout() time.Duration {
	return m.timeout
}

func (m *mockNetworkConfig) Retry() int {
	return m.retry
}

func (m *mockNetworkConfig) RetrySleepDuration() time.Duration {
	return m.sleep
}

type responseCatcher struct {
	Status string `json:"status"`
}

func TestNetwork(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	e := echo.New()

	e.GET("/200", func(e echo.Context) error {
		e.JSON(http.StatusOK, map[string]interface{}{"status": "200"})
		return nil
	})

	e.GET("/500", func(e echo.Context) error {
		e.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "500"})
		return nil
	})

	e.GET("/422", func(e echo.Context) error {
		e.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"status": "422"})
		return nil
	})

	e.POST("/200", func(e echo.Context) error {
		e.JSON(http.StatusOK, map[string]interface{}{"status": "200"})
		return nil
	})

	e.POST("/500", func(e echo.Context) error {
		e.JSON(http.StatusInternalServerError, map[string]interface{}{"status": "500"})
		return nil
	})

	go func() { e.Start(":6969") }()

	time.Sleep(time.Millisecond * 150)

	slinger := sling.New()

	ctx := context.Background()

	t.Run("GET", func(t *testing.T) {
		t.Run("When 200 the it will return nil and parse response in success binder", func(t *testing.T) {
			network := adapter.AdaptNetwork(slinger)

			cfg := &mockNetworkConfig{
				host:    "http://localhost:6969",
				timeout: time.Millisecond * 500,
				retry:   3,
				sleep:   time.Millisecond * 100,
			}

			response := responseCatcher{}
			assert.Nil(t, network.GET(ctx, cfg, "/200", &response, nil))
			assert.Equal(t, "200", response.Status)
		})

		t.Run("When >= 400 the it will return error and parse response in failed binder", func(t *testing.T) {
			network := adapter.AdaptNetwork(slinger)

			cfg := &mockNetworkConfig{
				host:    "http://localhost:6969",
				timeout: time.Millisecond * 500,
				retry:   3,
				sleep:   time.Millisecond * 100,
			}

			response := responseCatcher{}
			failedResponse := responseCatcher{}
			assert.NotNil(t, network.GET(ctx, cfg, "/422", &response, &failedResponse))
			assert.Equal(t, "", response.Status)
			assert.Equal(t, "422", failedResponse.Status)
		})

		t.Run("When >= 500 the it will return error and parse response in failed binder", func(t *testing.T) {
			network := adapter.AdaptNetwork(slinger)

			cfg := &mockNetworkConfig{
				host:    "http://localhost:6969",
				timeout: time.Millisecond * 500,
				retry:   3,
				sleep:   time.Millisecond * 100,
			}

			ctx := context.WithValue(ctx, "request-id", "abc")
			response := responseCatcher{}
			failedResponse := responseCatcher{}
			assert.NotNil(t, network.GET(ctx, cfg, "/500", &response, &failedResponse))
			assert.Equal(t, "", response.Status)
			assert.Equal(t, "500", failedResponse.Status)
		})
	})

	t.Run("POST", func(t *testing.T) {
		t.Run("When 200 the it will return nil and parse response in success binder", func(t *testing.T) {
			network := adapter.AdaptNetwork(slinger)

			cfg := &mockNetworkConfig{
				host:    "http://localhost:6969",
				timeout: time.Millisecond * 500,
				retry:   3,
				sleep:   time.Millisecond * 100,
			}

			response := responseCatcher{}
			assert.Nil(t, network.POST(ctx, cfg, "/200", nil, &response, nil))
			assert.Equal(t, "200", response.Status)
		})
	})

	_ = e.Close()
}
