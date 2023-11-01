package util

// constants for all supported user roles
const (
	SYSADMIN    = "system-admin"
	SUPER_ADMIN = "super-admin"
	ADMIN       = "admin"
	OPERATOR    = "operator"
)

func IsSupportedUserRole(role string) bool {
	switch role {
	case ADMIN, OPERATOR:
		return true
	}
	return false
}
