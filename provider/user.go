package provider

import (
	"context"

	"github.com/coronatorid/core-onator/entity"
)

//go:generate mockgen -source=./user.go -destination=./mocks/user_mock.go -package mockProvider

// User provider handling all scope about managing users
type User interface {
	Find(ctx context.Context, ID int) (entity.User, *entity.ApplicationError)

	// Phone number should be in phone number format
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, *entity.ApplicationError)

	Create(ctx context.Context, userInsertable entity.UserInsertable) (int, *entity.ApplicationError)
	CreateOrFind(ctx context.Context, phoneNumber string) (entity.User, *entity.ApplicationError)
}
