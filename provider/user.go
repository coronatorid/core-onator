package provider

import (
	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./user.go -destination=./mocks/user_mock.go -package mockProvider

// User provider handling all scope about managing users
type User interface {
	Find(ctx Context, ID int) (entity.User, *entity.ApplicationError)

	// Phone number should be in phone number format
	FindByPhoneNumber(ctx Context, phoneNumber string) (entity.User, *entity.ApplicationError)

	Create(ctx Context, userInsertable entity.UserInsertable) (int, *entity.ApplicationError)
	CreateOrFind(ctx Context, phoneNumber string) (entity.User, *entity.ApplicationError)
}
