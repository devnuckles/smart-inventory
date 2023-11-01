package util

// constants for all supported user status
const (
	PENDING  = "pending"
	APPROVED = "approved"
	REJECTED = "rejected"
)

func IsSupportedUserStatus(status string) bool {
	switch status {
	case PENDING, APPROVED, REJECTED:
		return true
	}
	return false
}
