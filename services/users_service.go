package services

import (
	"github.com/pprasha2/bookstore_users-api/domain/users"
	"github.com/pprasha2/bookstore_users-api/utils/errors"
)

//CreateUser - creates a user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	/*
		return nil, &errors.RestErr{
			Message: "Internal server error",
			Status:  http.StatusInternalServerError,
			Error:   "Internal Server Error",
		}
	*/
	if err := user.Validate(); err != nil {
		return &user, nil
	}
	return &user, nil
}
