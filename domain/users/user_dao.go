package users

import (
	"fmt"
	"github.com/ferza17/golang_bookstore-users-api/datasources/mysql/users_db"
	"github.com/ferza17/golang_bookstore-users-api/logger"
	"github.com/ferza17/golang_bookstore-users-api/utils/crypt"
	"github.com/ferza17/golang_bookstore-users-api/utils/date"
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
	mysql_utils "github.com/ferza17/golang_bookstore-users-api/utils/mysql"
	"strings"
)

const (
	queryInsertUser             = "INSERT INTO users(Firstname, Lastname,Email,DateCreated, Status, Password) VALUES (?,?,?,?,?,?);"
	queryGetUser                = "SELECT * FROM users WHERE ID=?"
	queryUpdateUser             = "UPDATE users SET Firstname=?, Lastname=?, Email=? , Status =?, Password=? WHERE ID=?"
	queryDeleteUser             = "DELETE FROM users WHERE ID=?"
	queryFindUserByStatus       = "SELECT ID,Firstname, Lastname,Email,DateCreated, Status FROM users WHERE Status =?"
	queryFindByEmailAndPassword = "SELECT ID, Firstname, Lastname, Email, DateCreated, Status FROM users WHERE Email=? AND Password=? AND Status =?"
)

// Communicate with database to GET DATA
func (user *User) Get() *errors.RestError {
	// ACCESS datasource users_db to communicate with database, GET query in constant variable
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database Error")
	}
	// Defer in compile after return, it's like stack. if u open connection u most to stop it to avoid memory leaks and secure the app
	defer stmt.Close()

	// executing query with ID that declare in constant and save result to result variable
	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status, &user.Password); getErr != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("Database Error")
	}
	return nil
}

// Communicate with database to SAVING DATA ( in case insert / create data to database )
func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("Database Error")
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	user.Status = StatusActive
	user.Password = crypt.GetMd5(user.Password)
	insertResult, err := stmt.Exec(&user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status, &user.Password)
	if err != nil {
		logger.Error("error when trying to execute user statement", err)
		return errors.NewInternalServerError("Database Error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last insert id after creating a new user", err)
		return errors.NewInternalServerError("Database Error")
	}
	user.ID = userId
	return nil
}

// Communicate with database to updating DATA
func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("Database Error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Firstname, user.Lastname, user.Email, user.Status, user.Password, user.ID)
	if err != nil {
		logger.Error("error when trying to execute user statement", err)
		return errors.NewInternalServerError("Database Error")
	}
	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("Database Error")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.ID); err != nil {
		logger.Error("error when trying to execute user statement", err)
		return errors.NewInternalServerError("Database Error")
	}

	return nil
}

func (user *User) Search(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return nil, errors.NewInternalServerError("Database Error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to querying status statement", err)
		return nil, errors.NewInternalServerError("Database Error")
	}
	defer rows.Close()

	// Check if slice of user
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying scan row ", err)
			return nil, errors.NewInternalServerError("Database Error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		logger.Error("error when user not found ", err)
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestError {
	// ACCESS datasource users_db to communicate with database, GET query in constant variable
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("Database Error")
	}
	// Defer in compile after return, it's like stack. if u open connection u most to stop it to avoid memory leaks and secure the app, it run before return statement
	defer stmt.Close()

	// executing query with email and password that declare in constant and save result to result variable
	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("Invalid user credentials")
		}
		logger.Error("error when trying to get user by email and password", err)
		return errors.NewInternalServerError("Database Error")
	}
	return nil
}
