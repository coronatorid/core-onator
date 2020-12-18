package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// RevokeToken that previously granted to application
type RevokeToken struct{}

// Perform grant token logic
func (g *RevokeToken) Perform(ctx provider.Context, request entity.RevokeTokenRequest, altairCfg provider.NetworkConfig, network provider.Network) *entity.ApplicationError {
	var altairError entity.AltairError

	encodedJSON, _ := json.Marshal(request)
	if err := network.POST(ctx, altairCfg, "/_plugins/oauth/authorizations/revoke", bytes.NewBuffer(encodedJSON), nil, &altairError); err != nil {
		// TODO: do proper error handling here
		fmt.Println("ERROR SENDING REQUEST", altairError)
		return err
	}

	return nil
}
