package services

import (
	"net/http"

	"github.com/pprasha2/bookstore_users-api/domain/users"
	"github.com/pprasha2/bookstore_users-api/utils/errors"
)

//CreateUser - creates a user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	//	return &user, nil
	return nil, &errors.RestErr{
		Message: "Internal server error",
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}
