package usecase_test

import (
	"testing"
	"time"

	"github.com/coronatorid/core-onator/provider/altair/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNetworkConfig(t *testing.T) {
	Host := "http://localhost:1304"
	Username := "altair"
	Password := "altair"
	Timeout := time.Second
	Retry := 3
	SleepDuration := time.Second

	cfg := &usecase.NetworkConfig{}
	cfg.Cfg.Host = Host
	cfg.Cfg.Username = Username
	cfg.Cfg.Password = Password
	cfg.Cfg.Timeout = Timeout
	cfg.Cfg.Retry = Retry
	cfg.Cfg.SleepDuration = SleepDuration

	assert.Equal(t, Host, cfg.Host())
	assert.Equal(t, Username, cfg.Username())
	assert.Equal(t, Password, cfg.Password())
	assert.Equal(t, Timeout, cfg.Timeout())
	assert.Equal(t, Retry, cfg.Retry())
	assert.Equal(t, SleepDuration, cfg.RetrySleepDuration())
}
