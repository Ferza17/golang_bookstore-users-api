package users

import (
	"fmt"
	"github.com/ferza17/golang_bookstore-users-api/util/errors"
)

var (
	usersDB = make(map[int64]*User)
)



func (user *User) Get()  *errors.RestError{
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = result.ID
	user.Firstname = result.Firstname
	user.Lastname = result.Lastname
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError  {
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists!", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists!", user.ID))
	}

	usersDB[user.ID] = user
	return nil
}