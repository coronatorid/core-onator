package auth

// Auth provide authorization and authentication for coronator
type Auth struct{}

// Fabricate auth service for coronator
func Fabricate() *Auth {
	return &Auth{}
}
