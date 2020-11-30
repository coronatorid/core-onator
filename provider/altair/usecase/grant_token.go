package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// GrantToken to altair oauth service
type GrantToken struct{}

// Perform grant token logic
func (g *GrantToken) Perform(ctx context.Context, request entity.GrantTokenRequest, altairCfg provider.NetworkConfig, network provider.Network) (entity.OauthAccessToken, *entity.ApplicationError) {
	var oauthAccessToken entity.OauthAccessToken
	var altairError entity.AltairError

	encodedJSON, _ := json.Marshal(request)
	if err := network.POST(ctx, altairCfg, "/_plugins/oauth/authorizations", bytes.NewBuffer(encodedJSON), &oauthAccessToken, &altairError); err != nil {
		// TODO: do proper error handling here
		fmt.Println("ERROR SENDING REQUEST", altairError)
		return oauthAccessToken, err
	}

	return oauthAccessToken, nil
}
