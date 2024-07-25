package database

import (
	model "assignment-totality-corp/internal/models"

	"github.com/stretchr/testify/mock"
)

// MockDatabase is a mock implementation of the InterfacesDatabase interface
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetUserById(id int32) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockDatabase) GetUserListByIds(ids []int32) ([]model.User, error) {
	args := m.Called(ids)
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockDatabase) GetAllUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}
