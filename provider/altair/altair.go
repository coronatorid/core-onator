package altair

import (
	"context"

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
func Fabricate(network provider.Network, altairNetworkCfg provider.NetworkConfig) (*Altair, error) {
	return &Altair{
		network:          network,
		altairNetworkCfg: altairNetworkCfg,
	}, nil
}

// GrantToken granting access token from altair
func (a *Altair) GrantToken(ctx context.Context, request entity.GrantTokenRequest) (entity.OauthAccessToken, *entity.ApplicationError) {
	grantToken := usecase.GrantToken{}
	return grantToken.Perform(ctx, request, a.altairNetworkCfg, a.network)
}
