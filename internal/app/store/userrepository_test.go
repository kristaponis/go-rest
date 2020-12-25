package store_test

import (
	"testing"

	"github.com/kristaponis/go-rest/internal/app/models"
	"github.com/kristaponis/go-rest/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	usr, err := s.User().Create(&models.User{
		ID: "1",
		Email: "user@example.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, usr)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&models.User{
		ID: "1",
		Email: "user@example.com",
	})
	usr, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, usr)
}