package user

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/user/usecase"
)

// User provide function for managing user data
type User struct {
	db provider.DB
}

// Fabricate user provider
func Fabricate(db provider.DB) *User {
	return &User{db: db}
}

// Find user by id
func (u *User) Find(ctx provider.Context, ID int) (entity.User, *entity.ApplicationError) {
	find := usecase.Find{}
	return find.Perform(ctx, ID, u.db)
}

// FindByPhoneNumber find user by phone number
func (u *User) FindByPhoneNumber(ctx provider.Context, phoneNumber string) (entity.User, *entity.ApplicationError) {
	findByPhoneNumber := usecase.FindByPhoneNumber{}
	return findByPhoneNumber.Perform(ctx, phoneNumber, u.db)
}

// Create user
func (u *User) Create(ctx provider.Context, userInsertable entity.UserInsertable) (int, *entity.ApplicationError) {
	create := usecase.Create{}
	return create.Perform(ctx, userInsertable, u.db)
}

// CreateOrFind user based on its phone number
func (u *User) CreateOrFind(ctx provider.Context, phoneNumber string) (entity.User, *entity.ApplicationError) {
	createOrFind := usecase.CreateOrFindUser{}
	return createOrFind.Perform(ctx, phoneNumber, u)
}
