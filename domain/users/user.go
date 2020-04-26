package users

import (
	"fmt"
	"strings"

	"github.com/pprasha2/bookstore_users-api/utils/errors"
)

//User structure
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate validates the checks before processing further
func (user *User) Validate() *errors.RestErr {
	fmt.Println("In validating Method")
	fmt.Println(user.Email)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		fmt.Println("Email is empty")
		return errors.NewBadRequestError("Invalid Email address")
	}
	fmt.Println("check again")
	fmt.Println(user.Email)
	return nil
}
