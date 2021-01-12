package admin

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/admin/api"
	"github.com/coronatorid/core-onator/provider/admin/usecase"
)

// Admin ...
type Admin struct {
	altair         provider.Altair
	authProvider   provider.Auth
	userProvider   provider.User
	reportProvider provider.Report
}

// Fabricate admin ...
func Fabricate(altair provider.Altair, authProvider provider.Auth, userProvider provider.User, reportProvider provider.Report) *Admin {
	return &Admin{
		altair:         altair,
		authProvider:   authProvider,
		userProvider:   userProvider,
		reportProvider: reportProvider,
	}
}

// FabricateAPI fabricating auth related API
func (a *Admin) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewAuthorizationRequestOTP(a))
	engine.InjectAPI(api.NewAuthorizationLogin(a))
	engine.InjectAPI(api.NewReportList(a))
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

// Authenticate admin
func (a *Admin) Authenticate(ctx provider.Context, userID int, allowedRole []constant.UserRole) (entity.User, *entity.ApplicationError) {
	authenticate := usecase.Authenticate{}
	return authenticate.Perform(ctx, userID, allowedRole, a.userProvider)
}

// ReportList ...
func (a *Admin) ReportList(ctx provider.Context, userID int, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta) ([]entity.ReportedCases, entity.ResponseMeta, *entity.ApplicationError) {
	reportList := usecase.ReportList{}
	return reportList.Perform(ctx, userID, status, requestMeta, a, a.reportProvider)
}
