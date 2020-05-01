package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/pprasha2/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		fmt.Println("in thissss")
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), errorNoRows) {
			fmt.Println("in in")
			return errors.NewNotFoundError("user not found with given id")
		}
		fmt.Println("in out")
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlError.Number {
	case 1062:
		fmt.Println(sqlError)
		return errors.NewBadRequestError("invalid data")

	}
	return errors.NewInternalServerError("error processing request")
}
