package constant

// UserRole ...
type UserRole int

const (
	// UserRoleNormal ...
	UserRoleNormal UserRole = iota
	// UserRoleSuperAdmin ...
	UserRoleSuperAdmin
)

// Int convert user role to integer
func (u UserRole) Int() int {
	return int(u)
}
