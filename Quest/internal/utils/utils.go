package utils

import "Quest/internal/types"

func ValidateUser(user types.User) bool {
	if user.Name == "" {
		return false
	}
	return true
}
