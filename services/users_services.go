package services

import (
	"github.com/ferza17/golang_bookstore-users-api/domain/users"
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	// VALIDATE user data
	if err := user.Validate(); err != nil {
		return nil, err
	}
	// Saving user to Database with user_dao (Data Access Object)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	// GET Reference of user and set ID
	result := &users.User{ID: userId}
	// Use user_dao to get Data from database and check if contains error
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	// use GetUser from this service that contain ID in database and check if contains error then return error
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	// Validate Request FROM user and Check if isPartial = false ( PUT METHOD ) then validate
	if !isPartial {
		if err := user.Validate(); err != nil {
			return nil, err
		}
	}

	// Updating USER domain with new Value from request user
	// Check if isPartial TRUE which means request is PATCH method ( Partial UPDATE )
	// else full updating value which means request is PUT method ( FULL UPDATE )
	if isPartial {
		if user.Firstname != "" {
			current.Firstname = user.Firstname
		}
		if user.Lastname != "" {
			current.Lastname = user.Lastname
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {

		current.Firstname = user.Firstname
		current.Lastname = user.Lastname
		current.Email = user.Email
	}

	// Access user_dao ( in case variable current is contain memory of user_dao ) and update value
	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *errors.RestError {
	// check if id exists
	if _, err := GetUser(userId); err != nil {
		return err
	}

	user := &users.User{ID: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestError) {
	dao := &users.User{}
	return dao.Search(status)
}
