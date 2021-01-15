package provider

//go:generate mockgen -source=./admin.go -destination=./mocks/admin_mock.go -package mockProvider

import (
	"github.com/coronatorid/core-onator/constant"
	"github.com/coronatorid/core-onator/entity"
)

// Admin provide function for administrators
type Admin interface {
	Login(ctx Context, request entity.Login) (entity.LoginResponse, *entity.ApplicationError)
	RequestOTP(ctx Context, request entity.RequestOTP) (*entity.RequestOTPResponse, *entity.ApplicationError)
	Authenticate(ctx Context, adminID int, allowedRole []constant.UserRole) (entity.User, *entity.ApplicationError)

	ReportList(ctx Context, adminID int, status constant.ReportedCasesStatus, requestMeta entity.RequestMeta) ([]entity.ReportedCases, entity.ResponseMeta, *entity.ApplicationError)
	ReportDelete(ctx Context, adminID int, reportedCasesID int) *entity.ApplicationError
	ReportReject(ctx Context, adminID, reportedCasesID int) (entity.ReportedCases, *entity.ApplicationError)
	ReportConfirm(ctx Context, adminID, reportedCasesID int) (entity.ReportedCases, *entity.ApplicationError)
}
