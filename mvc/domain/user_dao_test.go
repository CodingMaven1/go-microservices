package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "Did not expect a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 was not found!", err.Message)
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user, "Expected to return a user with id 123")
	assert.Nil(t, err)

	assert.EqualValues(t, 1, user.Id)
}
