package usecase

import (
	"errors"
	"net/http"

	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
)

// Authenticate admin role
type Authenticate struct{}

// Perform authenticate business logic
func (a *Authenticate) Perform(ctx provider.Context, adminID int, allowedRole []constant.UserRole, userProvider provider.User) (entity.User, *entity.ApplicationError) {
	user, err := userProvider.Find(ctx, adminID)
	if err != nil && err.HTTPStatus == http.StatusNotFound {
		return user, &entity.ApplicationError{
			Err:        []error{errors.New("only admin can use this feature")},
			HTTPStatus: http.StatusForbidden,
		}
	} else if err != nil {
		return user, err
	}

	forbidden := true
	for _, role := range allowedRole {
		if role.Int() == user.Role {
			forbidden = false
		}
	}

	if forbidden {
		return user, &entity.ApplicationError{
			Err:        []error{errors.New("only admin can use this feature")},
			HTTPStatus: http.StatusForbidden,
		}
	}

	return user, nil
}
