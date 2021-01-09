package admin

import (
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/admin/api"
	"github.com/coronatorid/core-onator/provider/admin/usecase"
)

// Admin ...
type Admin struct {
	altair       provider.Altair
	authProvider provider.Auth
	userProvider provider.User
}

// Fabricate admin ...
func Fabricate(altair provider.Altair, authProvider provider.Auth, userProvider provider.User) *Admin {
	return &Admin{
		altair:       altair,
		authProvider: authProvider,
		userProvider: userProvider,
	}
}

// FabricateAPI fabricating auth related API
func (a *Admin) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewRequestOTP(a))
	engine.InjectAPI(api.NewLogin(a))
}

// Login with admin flow
func (a *Admin) Login(ctx provider.Context, request entity.Login) (entity.LoginResponse, *entity.ApplicationError) {
	login := usecase.Login{}
	return login.Perform(ctx, request, a.userProvider, a.altair, a.authProvider)
}

// RequestOTP with admin flow
func (a *Admin) RequestOTP(ctx provider.Context, request entity.RequestOTP) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	requestOTP := usecase.RequestOTP{}
	return requestOTP.Perform(ctx, request, a.userProvider, a.authProvider)
}
