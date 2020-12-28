package users

import "github.com/ferza17/golang_bookstore-users-api/util/errors"

type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}



func (user *User) Validate()  *errors.RestError{
	if user.Email == "" {
		return  errors.NewBadRequestError("Invalid Email Address!")
	}
	return nil
}