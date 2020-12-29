package mysql_utils

import (
	"github.com/ferza17/golang_bookstore-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("Error parsing mysql response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid Data")
	}
	return errors.NewInternalServerError("Error processing request!")
}