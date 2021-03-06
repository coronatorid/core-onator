package auth

import (
	"os"
	"regexp"
	"time"

	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/auth/api"
	"github.com/coronatorid/core-onator/provider/auth/inappcron"
	"github.com/coronatorid/core-onator/provider/auth/usecase"

	"github.com/coronatorid/core-onator/entity"
)

// Auth provide authorization and authentication for coronator
type Auth struct {
	regexIndonesianPhoneNumber *regexp.Regexp
	otpRetryDuration           time.Duration

	cache                       provider.Cache
	textPublisher               provider.TextPublisher
	userProvider                provider.User
	altair                      provider.Altair
	whatsappPublisherFabricator func() (provider.TextPublisher, error)
}

// Fabricate auth service for coronator
func Fabricate(cache provider.Cache, textPublisher provider.TextPublisher, userProvider provider.User, altair provider.Altair, whatsappPublisherFabricator func() (provider.TextPublisher, error)) (*Auth, error) {
	regexIndonesianPhoneNumber, _ := regexp.Compile(`^\+62\d{10,12}`)
	d, err := time.ParseDuration(os.Getenv("OTP_RETRY_DURATION"))
	if err != nil {
		return nil, err
	}

	return &Auth{
		regexIndonesianPhoneNumber: regexIndonesianPhoneNumber,
		otpRetryDuration:           d,

		cache:                       cache,
		textPublisher:               textPublisher,
		userProvider:                userProvider,
		altair:                      altair,
		whatsappPublisherFabricator: whatsappPublisherFabricator,
	}, nil
}

// FabricateAPI fabricating auth related API
func (a *Auth) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewRequestOTP(a))
	engine.InjectAPI(api.NewLogin(a))
	engine.InjectAPI(api.NewLogout(a))
}

// FabricateInAppCronjob fabricating inappcronjob related process
func (a *Auth) FabricateInAppCronjob(cron provider.InAppCron) {
	cron.Inject(inappcron.NewRenewTextPublisher(a, a.whatsappPublisherFabricator))
}

// RequestOTP send otp based on request by the client
func (a *Auth) RequestOTP(ctx provider.Context, request entity.RequestOTP, otpDigit int) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	requestOTP := &usecase.RequestOTP{}
	return requestOTP.Perform(ctx, request, a.regexIndonesianPhoneNumber, a.cache, a.textPublisher, a.otpRetryDuration, otpDigit)
}

// Login ...
func (a *Auth) Login(ctx provider.Context, request entity.Login, otpDigit int) (entity.LoginResponse, *entity.ApplicationError) {
	login := &usecase.Login{}
	return login.Perform(ctx, request, a.otpRetryDuration, a.userProvider, a.altair, otpDigit, a)
}

// Logout ...
func (a *Auth) Logout(ctx provider.Context, request entity.RevokeTokenRequest) *entity.ApplicationError {
	logout := &usecase.Logout{}
	return logout.Perform(ctx, request, a.altair)
}

// RenewTextPublisher session
func (a *Auth) RenewTextPublisher(textPublisher provider.TextPublisher) {
	a.textPublisher = textPublisher
}

// ValidateOTP ...
func (a *Auth) ValidateOTP(ctx provider.Context, request entity.Login, otpDigit int) *entity.ApplicationError {
	validateOTP := usecase.ValidateOTP{}
	return validateOTP.Perform(ctx, request, a.otpRetryDuration, otpDigit)
}
