package provider

import (
	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./altair.go -destination=./mocks/altair_mock.go -package mockProvider

// Altair interface connector
type Altair interface {
	GrantToken(ctx Context, request entity.GrantTokenRequest) (entity.OauthAccessToken, *entity.ApplicationError)
	RevokeToken(ctx Context, request entity.RevokeTokenRequest) *entity.ApplicationError
}
