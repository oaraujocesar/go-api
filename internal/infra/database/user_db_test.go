package database

import (
	"testing"

	"github.com/oaraujocesar/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createInMemoryDatabase(t *testing.T) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&entity.User{})

	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := createInMemoryDatabase(t)

	user, _ := entity.NewUser("Cesar", "test@mail.com", "12341234")
	assert.Nil(t, err)

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.NotEmpty(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := createInMemoryDatabase(t)

	user, _ := entity.NewUser("Cesar", "test@mail.com", "12341234")
	assert.Nil(t, err)

	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.NotEmpty(t, userFound)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmailWhenUserDoesNotExists(t *testing.T) {
	db, err := createInMemoryDatabase(t)
	assert.Nil(t, err)

	userDB := NewUser(db)

	_, err = userDB.FindByEmail("notFound@mail.com")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "record not found")
}
