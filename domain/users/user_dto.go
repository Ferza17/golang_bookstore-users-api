package users

import (
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
)

type User struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

const (
	StatusActive = "active"
)

func (user *User) Validate() *errors.RestError {
	// Remove Space of Firstname
	//user.Firstname = strings.TrimSpace(user.Firstname)
	// Remove Space of Firstname
	//user.Lastname = strings.TrimSpace(user.Lastname)

	// Validate email if the value is empty
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address!")
	}
	return nil
}
