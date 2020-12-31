package users

import (
	"fmt"
	"github.com/ferza17/golang_bookstore-users-api/datasources/mysql/users_db"
	"github.com/ferza17/golang_bookstore-users-api/utils/crypt"
	"github.com/ferza17/golang_bookstore-users-api/utils/date"
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
	mysqlUtils "github.com/ferza17/golang_bookstore-users-api/utils/mysql"
)

const (
	queryInsertUser       = "INSERT INTO users(Firstname, Lastname,Email,DateCreated, Status, Password) VALUES (?,?,?,?,?,?);"
	queryGetUser          = "SELECT * FROM users WHERE ID=?"
	queryUpdateUser       = "UPDATE users SET Firstname=?, Lastname=?, Email=? , Status =?, Password=? WHERE ID=?"
	queryDeleteUser       = "DELETE FROM users WHERE ID=?"
	queryFindUserByStatus = "SELECT ID,Firstname, Lastname,Email,DateCreated, Status FROM users WHERE Status =?"
)

// Communicate with database to GET DATA
func (user *User) Get() *errors.RestError {
	// ACCESS datasource users_db to communicate with database, GET query in constant variable
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// Defer in compile after return, it's like stack. if u open connection u most to stop it to avoid memory leaks and secure the app
	defer stmt.Close()

	// executing query with ID that declare in constant and save result to result variable
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status, &user.Password); getErr != nil {
		return mysqlUtils.ParseError(getErr)
	}
	return nil
}

// Communicate with database to SAVING DATA ( in case insert / create data to database )
func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	user.Status = StatusActive
	user.Password = crypt.GetMd5(user.Password)
	insertResult, err := stmt.Exec(&user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status, &user.Password)
	if err != nil {
		return mysqlUtils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlUtils.ParseError(err)
	}
	user.ID = userId
	return nil
}

// Communicate with database to updating DATA
func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Firstname, user.Lastname, user.Email, user.Status, user.Password, user.ID)
	if err != nil {
		return mysqlUtils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.ID); err != nil {
		return mysqlUtils.ParseError(err)
	}

	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	// Check if slice of user
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlUtils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
