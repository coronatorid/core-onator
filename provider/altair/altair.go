package altair

import (
	"os"
	"strconv"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/altair/usecase"
)

// Altair plugin connector
type Altair struct {
	network          provider.Network
	altairNetworkCfg provider.NetworkConfig
}

// Fabricate altair plugin connector
func Fabricate(network provider.Network) (*Altair, error) {
	networkCfg := &usecase.NetworkConfig{}

	networkCfg.Cfg.Host = os.Getenv("ALTAIR_HOST")
	networkCfg.Cfg.Username = os.Getenv("ALTAIR_BASIC_USERNAME")
	networkCfg.Cfg.Password = os.Getenv("ALTAIR_BASIC_PASSWORD")

	retryCount, err := strconv.Atoi(os.Getenv("ALTAIR_RETRY_COUNT"))
	if err != nil {
		return nil, err
	}
	networkCfg.Cfg.Retry = retryCount

	retrySleepDuration, err := time.ParseDuration(os.Getenv("ALTAIR_RETRY_SLEEP_DURATION"))
	if err != nil {
		return nil, err
	}
	networkCfg.Cfg.SleepDuration = retrySleepDuration

	return &Altair{
		network:          network,
		altairNetworkCfg: networkCfg,
	}, nil
}

// GrantToken granting access token from altair
func (a *Altair) GrantToken(ctx provider.Context, request entity.GrantTokenRequest) (entity.OauthAccessToken, *entity.ApplicationError) {
	grantToken := usecase.GrantToken{}
	return grantToken.Perform(ctx, request, a.altairNetworkCfg, a.network)
}
