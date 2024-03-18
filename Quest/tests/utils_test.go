package test

import (
	"testing"

	"Quest/internal/types"
	"Quest/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	validUser := types.User{Name: "Danila", Balance: 100}
	invalidUser := types.User{Name: "", Balance: 0}

	assert.True(t, utils.ValidateUser(validUser), "Expected valid user, got invalid")
	assert.False(t, utils.ValidateUser(invalidUser), "Expected invalid user, got valid")
}

func TestValidateQuest(t *testing.T) {
	validQuest := types.Quest{Name: "Print something", Cost: 50}
	invalidQuest := types.Quest{Name: "", Cost: 0}

	assert.True(t, utils.ValidateQuest(validQuest), "Expected valid quest, got invalid")
	assert.False(t, utils.ValidateQuest(invalidQuest), "Expected invalid quest, got valid")
}

func TestValidateSignal(t *testing.T) {
	validSignal := types.Signal{UserId: 1, QuestId: 2}
	invalidSignal := types.Signal{UserId: 0, QuestId: 0}

	assert.True(t, utils.ValidateSignal(validSignal), "Expected valid signal, got invalid")
	assert.False(t, utils.ValidateSignal(invalidSignal), "Expected invalid signal, got valid")
}
