package database

import (
	model "assignment-totality-corp/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBCreation(t *testing.T) {
	db := DBCreation()
	assert.Equal(t, 4, len(db.Users), "Expected 4 users in the database")
}

func TestGetUser(t *testing.T) {
	db := DBCreation()

	// Test retrieving an existing user
	user, err := db.GetUserById(1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), user.ID, "Expected user ID to be 1")

	// Test retrieving a non-existing user
	user, err = db.GetUserById(99)
	assert.Error(t, err, "Expected an error for non-existing user")
	assert.Equal(t, model.User{}, user, "Expected empty user for non-existing user")
}

func TestGetUserList(t *testing.T) {
	db := DBCreation()

	// Test retrieving a list of users
	users, err := db.GetUserListByIds([]int32{1, 2, 99})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users), "Expected 2 users in the result")

	// Verify that the correct users are returned
	expectedIDs := map[int32]bool{1: true, 2: true}
	for _, user := range users {
		assert.True(t, expectedIDs[user.ID], "Expected user ID to be 1 or 2")
	}
}

func TestGetUsers(t *testing.T) {
	db := DBCreation()

	// Retrieve the list of users
	users, err := db.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(users), "Expected 4 users in the result")

	// Verify that the retrieved users match the database users
	ids := make(map[int32]bool)
	for _, user := range users {
		ids[user.ID] = true
	}
	for id := range db.Users {
		assert.True(t, ids[id], "Expected user ID to be present in the result")
	}
}
