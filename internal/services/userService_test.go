package services

import (
	"errors"
	"testing"

	model "assignment-totality-corp/internal/models"
	mocks "assignment-totality-corp/mocks/database"

	"github.com/stretchr/testify/assert"
)

// Tests for UserService methods

func TestGetUserById(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	service := NewUserService(mockDB)

	mockDB.On("GetUser", int32(1)).Return(model.User{ID: 1, FullName: "Alice"}, nil)

	user, err := service.GetUserById(1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), user.ID)
	assert.Equal(t, "Alice", user.FullName)

	mockDB.On("GetUser", int32(99)).Return(model.User{}, errors.New("user not found"))

	user, err = service.GetUserById(99)
	assert.Error(t, err)
	assert.Equal(t, model.User{}, user)
}

func TestGetUserByIds(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	service := NewUserService(mockDB)

	mockDB.On("GetUserList", []int32{1, 2, 99}).Return([]model.User{
		{ID: 1, FullName: "Alice"},
		{ID: 2, FullName: "Bob"},
	}, nil)

	users, err := service.GetUserByIds([]int32{1, 2, 99})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))
	assert.Contains(t, users, model.User{ID: 1, FullName: "Alice"})
	assert.Contains(t, users, model.User{ID: 2, FullName: "Bob"})
}

func TestSearchUsers(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	service := NewUserService(mockDB)

	mockDB.On("GetUsers").Return([]model.User{
		{ID: 1, FullName: "Alice", City: "New York", Phone: 1234567890, Height: 5.5, Married: false},
		{ID: 2, FullName: "Bob", City: "Los Angeles", Phone: 1234567890, Height: 5.5, Married: true},
		{ID: 3, FullName: "Charlie", City: "Chicago", Phone: 1234567890, Height: 5.5, Married: false},
		{ID: 4, FullName: "David", City: "New York", Phone: 1234567890, Height: 5.5, Married: true},
	}, nil)

	tests := []struct {
		name     string
		request  model.SearchUsersRequest
		expected []model.User
	}{
		{
			name:    "Search by FullName",
			request: model.SearchUsersRequest{Fname: "Alice"},
			expected: []model.User{
				{ID: 1, FullName: "Alice", City: "New York", Phone: 1234567890, Height: 5.5, Married: false},
			},
		},
		{
			name:    "Search by City",
			request: model.SearchUsersRequest{City: "New York"},
			expected: []model.User{
				{ID: 1, FullName: "Alice", City: "New York", Phone: 1234567890, Height: 5.5, Married: false},
				{ID: 4, FullName: "David", City: "New York", Phone: 1234567890, Height: 5.5, Married: true},
			},
		},
		{
			name:    "Search by MinHeight",
			request: model.SearchUsersRequest{MinHeight: 5.5},
			expected: []model.User{
				{ID: 1, FullName: "Alice", City: "New York", Phone: 1234567890, Height: 5.5, Married: false},
				{ID: 2, FullName: "Bob", City: "Los Angeles", Phone: 1234567890, Height: 5.5, Married: true},
				{ID: 3, FullName: "Charlie", City: "Chicago", Phone: 1234567890, Height: 5.5, Married: false},
				{ID: 4, FullName: "David", City: "New York", Phone: 1234567890, Height: 5.5, Married: true},
			},
		},
		{
			name:    "Search by Married",
			request: model.SearchUsersRequest{Married: Bool(true)},
			expected: []model.User{
				{ID: 2, FullName: "Bob", City: "Los Angeles", Phone: 1234567890, Height: 5.5, Married: true},
				{ID: 4, FullName: "David", City: "New York", Phone: 1234567890, Height: 5.5, Married: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := service.SearchUsers(tt.request)
			assert.NoError(t, err)
			assert.ElementsMatch(t, tt.expected, users)
		})
	}

	mockDB.On("GetUsers").Return(nil, assert.AnError)
	users, err := service.SearchUsers(model.SearchUsersRequest{})
	assert.NoError(t, err)
	assert.Equal(t, len(users), 0)
}

func Bool(b bool) *bool {
	return &b
}
