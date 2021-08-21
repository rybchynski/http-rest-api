package store_test

import (
	"testing"

	"github.com/rybchynski/http-rest-api/internal/app/model"
	"github.com/rybchynski/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user2@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user2@example.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user3@example.com",
	})
	email2 := "user3@example.com"
	u, err := s.User().FindByEmail(email2)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
