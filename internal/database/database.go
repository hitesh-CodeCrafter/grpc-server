package database

import (
	model "assignment-totality-corp/internal/models"
	"errors"
	"log"
)

type InterfaceDb interface {
	GetUserById(int32) (model.User, error)
	GetUserListByIds([]int32) ([]model.User, error)
	GetAllUsers() ([]model.User, error)
}

type Database struct {
	Users map[int32]model.User
}

func DBCreation() Database {
	users := make(map[int32]model.User)

	users[1] = model.User{
		ID:       1,
		FullName: "SHIVAM",
		City:     "LA",
		Phone:    1234567890,
		Height:   4,
		Married:  true,
	}

	users[2] = model.User{
		ID:       2,
		FullName: "Ben",
		City:     "Los Angeles",
		Phone:    0,
		Height:   4.3,
		Married:  true,
	}

	return Database{Users: users}
}

func (db *Database) GetUserById(id int32) (model.User, error) {
	user, ok := db.Users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (db *Database) GetUserListByIds(ids []int32) ([]model.User, error) {
	users := make([]model.User, 0)
	for _, id := range ids {
		user, ok := db.Users[id]
		if !ok {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *Database) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	for _, user := range db.Users {
		users = append(users, user)
	}
	log.Println(users, "users-->")
	return users, nil
}
