package auth

// Auth provide authorization and authentication for coronator
type Auth struct{}

// Fabricate auth service for coronator
func Fabricate() *Auth {
	return &Auth{}
}

// RequestOTP send otp based on request by the client
func (a *Auth) RequestOTP() {

}
