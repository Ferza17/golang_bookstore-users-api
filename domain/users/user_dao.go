package users

import (
	"fmt"
	"github.com/ferza17/golang_bookstore-users-api/datasources/mysql/users_db"
	"github.com/ferza17/golang_bookstore-users-api/utils/date"
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
	"strings"
)


const (
	indexUniqueEmail = "Email"
	queryInsertUser = "INSERT INTO users(Firstname, Lastname,Email,DateCreated) VALUES (?,?,?,?);"
)




func (user *User) Get()  *errors.RestError{
	//if err := users_db.Client.Ping(); err != nil {
	//	panic(err)
	//}
	//
	//result := usersDB[user.ID]
	//if result == nil {
	//	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	//}
	//
	//user.ID = result.ID
	//user.Firstname = result.Firstname
	//user.Lastname = result.Lastname
	//user.Email = result.Email
	//user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError  {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()

	insertResult, err := stmt.Exec(user.Firstname, user.Lastname, user.Email, user.DateCreated)
	if err != nil {
		if  strings.Contains(err.Error(), indexUniqueEmail){
			return errors.NewBadRequestError(fmt.Sprintf("Email %s Already exists!", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error When trying to save user : %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error When trying to save user : %s", err.Error()))
	}
	user.ID = userId
	return nil
}