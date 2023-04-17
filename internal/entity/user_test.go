package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("César", "test@mail.com", "test123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "César", user.Name)
	assert.Equal(t, "test@mail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("César", "test@mail.com", "test123")
	assert.Nil(t, err)

	assert.True(t, user.ValidatePassword("test123"))
	assert.False(t, user.ValidatePassword("test"))
	assert.False(t, user.ValidatePassword("test1234"))
	assert.NotEqual(t, "test123", user.Password)
}
