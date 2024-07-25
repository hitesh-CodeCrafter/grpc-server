package services

import (
	"assignment-totality-corp/internal/database"
	model "assignment-totality-corp/internal/models"
	"log"
)

type InterfaceUserService interface {
	GetUserById(int32) (model.User, error)
	GetUserByIds([]int32) ([]model.User, error)
	SearchUsers(model.SearchUsersRequest) ([]model.User, error)
}

type UserService struct {
	db database.InterfaceDb
}

func NewUserService(db database.InterfaceDb) InterfaceUserService {
	return &UserService{db: db}
}

func (us *UserService) SearchUsers(req model.SearchUsersRequest) ([]model.User, error) {
	if req.Fname == "" && req.City == "" && req.Phone == 0 && req.MinHeight == 0 && req.MaxHeight == 0 && req.Married == nil {
		return []model.User{}, nil
	}

	users, err := us.db.GetAllUsers()
	if err != nil {
		return nil, err
	}

	log.Println("req", req)
	log.Println("users---->", users)

	filteredUsers := make([]model.User, 0)

	// filter users based on search criteria
	for i := 0; i < len(users); i++ {
		if req.Fname != "" && users[i].FullName != req.Fname {
			continue
		}
		if req.City != "" && users[i].City != req.City {
			continue
		}
		if req.Phone != 0 && users[i].Phone != req.Phone {
			continue
		}
		if req.MinHeight != 0 && users[i].Height < req.MinHeight {
			continue
		}
		if req.MaxHeight != 0 && users[i].Height > req.MaxHeight {
			continue
		}
		if req.Married != nil && users[i].Married != *req.Married {
			continue
		}

		filteredUsers = append(filteredUsers, users[i])
	}

	return filteredUsers, nil
}

func (us *UserService) GetUserById(id int32) (model.User, error) {
	user, err := us.db.GetUserById(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *UserService) GetUserByIds(ids []int32) ([]model.User, error) {
	users, err := us.db.GetUserListByIds(ids)
	if err != nil {
		return nil, err
	}
	return users, nil
}
