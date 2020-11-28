package usecase

import (
	"context"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// CreateOrFindUser create or find user based on given phone number
type CreateOrFindUser struct{}

// Perform create or find user logic
func (c *CreateOrFindUser) Perform(ctx context.Context, phoneNumber string, userProvider provider.User) (entity.User, *entity.ApplicationError) {
	user, err := userProvider.FindByPhoneNumber(ctx, phoneNumber)
	if err != nil && err.Err[0].Error() == "user not found" {
		id, err := userProvider.Create(ctx, entity.UserInsertable{PhoneNumber: phoneNumber})
		if err != nil {
			return user, err
		}

		user, err := userProvider.Find(ctx, id)
		if err != nil {
			return user, err
		}

		return user, nil
	} else if err != nil {
		return user, err
	}

	return user, nil
}
