package users

import (
	"fmt"

	"github.com/pprasha2/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//Get fetch user from DB
func (user *User) Get() *errors.RestErr{
result:= userDB[user.Id]
if result ==nil{
	return errors.NewNotFoundError(fmt.Sprintf("User %d not found ",user.Id))
}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName  = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil

}
//Save interacts with the databas e
func (user User) Save() *errors.RestErr{
	current := userDB[user.Id]
	if  current != nil{
		if current.Email == user.Email{
			return errors.NewBadRequestError(fmt.Sprintf("EmailId %s already registered!",user.Email)) 
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists!",user.Id)) 
	}
	userDB[user.Id] = &user
	return nil
}
