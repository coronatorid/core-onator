package provider

import (
	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./auth.go -destination=./mocks/auth_mock.go -package mockProvider

// Auth provider handle all authorization and authentication domain
type Auth interface {
	Login(ctx Context, request entity.Login, otpDigit int) (entity.LoginResponse, *entity.ApplicationError)
	Logout(ctx Context, request entity.RevokeTokenRequest) *entity.ApplicationError
	RequestOTP(ctx Context, request entity.RequestOTP, otpDigit int) (*entity.RequestOTPResponse, *entity.ApplicationError)
	RenewTextPublisher(textPublisher TextPublisher)
}
