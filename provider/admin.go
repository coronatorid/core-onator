package provider

//go:generate mockgen -source=./admin.go -destination=./mocks/admin_mock.go -package mockProvider

import "github.com/coronatorid/core-onator/entity"

// Admin provide function for administrators
type Admin interface {
	Login(ctx Context, request entity.Login) (entity.LoginResponse, *entity.ApplicationError)
	RequestOTP(ctx Context, request entity.RequestOTP) (*entity.RequestOTPResponse, *entity.ApplicationError)
}
