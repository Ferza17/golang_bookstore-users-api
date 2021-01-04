package users

import (
	"encoding/json"
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
)

type PublicUser struct {
	ID          int64  `json:"user_id"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
}
type PrivateUser struct {
	ID          int64  `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
	Status      string `json:"status"`
}

func (users Users) Marshall (isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users{
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall (isPublic bool) interface{}{
	// Approach 1
	if isPublic{
		return PublicUser{
			ID: user.ID,
			DateCreated: user.DateCreated,
			Status: user.Status,
		}
	}
	// Approach 2
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
	 	return errors.NewInternalServerError("error when unmarshalling private user")
	}
	return privateUser

}