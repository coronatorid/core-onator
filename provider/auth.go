package provider

import (
	"context"

	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./auth.go -destination=./mocks/auth_mock.go -package mockProvider

// Auth provider handle all authorization and authentication domain
type Auth interface {
	Login(ctx context.Context, request entity.Login) (entity.LoginResponse, *entity.ApplicationError)
	RequestOTP(ctx context.Context, request entity.RequestOTP) (*entity.RequestOTPResponse, *entity.ApplicationError)
}
