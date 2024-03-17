package utils

import "Quest/internal/types"

func ValidateUser(user types.User) bool {
	if user.Name == "" {
		return false
	}
	return true
}

func ValidateQuest(quest types.Quest) bool {
	if quest.Name == "" || quest.Cost == 0 {
		return false
	}
	return true
}
