package server

import (
	"context"
	"testing"

	model "assignment-totality-corp/internal/models"
	mocks "assignment-totality-corp/mocks/services"
	pb "assignment-totality-corp/proto/userService"

	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	us := NewUserService(mockUserService)

	user := model.User{
		ID:       1,
		FullName: "Alice",
		City:     "New York",
		Phone:    1234567890,
		Height:   5.5,
		Married:  false,
	}

	mockUserService.On("GetUserById", int32(1)).Return(user)

	req := &pb.GetUserRequest{Id: 1}
	resp, err := us.(*userService).RetrieveUserById(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, user.ID, resp.Id)
	assert.Equal(t, user.FullName, resp.FullName)
	assert.Equal(t, user.City, resp.City)
	assert.Equal(t, user.Phone, resp.Phone)
	assert.Equal(t, user.Height, resp.Height)
	assert.Equal(t, user.Married, resp.IsMarried)

	mockUserService.AssertExpectations(t)
}

func TestGetUsersByIds(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	us := NewUserService(mockUserService)

	users := []model.User{
		{
			ID:       1,
			FullName: "Alice",
			City:     "New York",
			Phone:    1234567890,
			Height:   5.5,
			Married:  false,
		},
		{
			ID:       2,
			FullName: "Bob",
			City:     "Los Angeles",
			Phone:    1234567890,
			Height:   5.5,
			Married:  true,
		},
	}

	mockUserService.On("GetUserByIds", []int32{1, 2}).Return(users)

	req := &pb.GetUsersRequest{Ids: []int32{1, 2}}
	resp, err := us.(*userService).RetrieveUsersByIds(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, len(users), len(resp.Users))

	for i, user := range users {
		assert.Equal(t, user.ID, resp.Users[i].Id)
		assert.Equal(t, user.FullName, resp.Users[i].FullName)
		assert.Equal(t, user.City, resp.Users[i].City)
		assert.Equal(t, user.Phone, resp.Users[i].Phone)
		assert.Equal(t, user.Height, resp.Users[i].Height)
		assert.Equal(t, user.Married, resp.Users[i].IsMarried)
	}

	mockUserService.AssertExpectations(t)
}

func TestSearchUsers(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	us := NewUserService(mockUserService)

	users := []model.User{
		{
			ID:       1,
			FullName: "Alice",
			City:     "New York",
			Phone:    1234567890,
			Height:   5.5,
			Married:  false,
		},
		{
			ID:       2,
			FullName: "Bob",
			City:     "Los Angeles",
			Phone:    1234567890,
			Height:   5.5,
			Married:  true,
		},
	}

	searchReq := model.SearchUsersRequest{
		Fname:     "Alice",
		City:      "New York",
		MinHeight: 0,
		MaxHeight: 10,
	}

	mockUserService.On("SearchUsers", searchReq).Return(users)

	req := &pb.SearchUsersRequest{
		FullName:  "Alice",
		City:      "New York",
		MinHeight: 0,
		MaxHeight: 10,
	}

	resp, err := us.(*userService).SearchUsers(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, len(users), len(resp.Users))

	for i, user := range users {
		assert.Equal(t, user.ID, resp.Users[i].Id)
		assert.Equal(t, user.FullName, resp.Users[i].FullName)
		assert.Equal(t, user.City, resp.Users[i].City)
		assert.Equal(t, user.Phone, resp.Users[i].Phone)
		assert.Equal(t, user.Height, resp.Users[i].Height)
		assert.Equal(t, user.Married, resp.Users[i].IsMarried)
	}

	mockUserService.AssertExpectations(t)
}
