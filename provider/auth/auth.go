package auth

import (
	"context"
	"os"
	"regexp"
	"time"

	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/provider/auth/api"
	"github.com/coronatorid/core-onator/provider/auth/usecase"

	"github.com/coronatorid/core-onator/entity"
)

// Auth provide authorization and authentication for coronator
type Auth struct {
	regexIndonesianPhoneNumber *regexp.Regexp
	otpRetryDuration           time.Duration

	cache         provider.Cache
	textPublisher provider.TextPublisher
	userProvider  provider.User
}

// Fabricate auth service for coronator
func Fabricate(cache provider.Cache, textPublisher provider.TextPublisher, userProvider provider.User) (*Auth, error) {
	regexIndonesianPhoneNumber, _ := regexp.Compile(`^\+62\d{10,12}`)
	d, err := time.ParseDuration(os.Getenv("OTP_RETRY_DURATION"))
	if err != nil {
		return nil, err
	}

	return &Auth{
		regexIndonesianPhoneNumber: regexIndonesianPhoneNumber,
		otpRetryDuration:           d,

		cache:         cache,
		textPublisher: textPublisher,
		userProvider:  userProvider,
	}, nil
}

// FabricateAPI fabricating auth related API
func (a *Auth) FabricateAPI(engine provider.APIEngine) {
	engine.InjectAPI(api.NewRequestOTP(a))
}

// RequestOTP send otp based on request by the client
func (a *Auth) RequestOTP(ctx context.Context, request entity.RequestOTP) (*entity.RequestOTPResponse, *entity.ApplicationError) {
	requestOTP := &usecase.RequestOTP{}
	return requestOTP.Perform(ctx, request, a.regexIndonesianPhoneNumber, a.cache, a.textPublisher, a.otpRetryDuration)
}
