package users

import (
	"fmt"

	"github.com/pprasha2/bookstore_users-api/datasources/mysql/users_db"
	"github.com/pprasha2/bookstore_users-api/utils/date_utils"
	"github.com/pprasha2/bookstore_users-api/utils/errors"
	mysql_utils "github.com/pprasha2/bookstore_users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id = ?"
	errorNoRows      = "no rows in result set"
)

var (
	userDB = make(map[int64]*User)
)

//dao --> data access object

//Get fetch user from DB
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError("Internal server error")
	}
	defer stmt.Close()
	fmt.Println("ERRrrrrrrr")
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil

}

//Save interacts with the databas e
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("Internal server error")
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if SaveErr != nil {
		return mysql_utils.ParseError(SaveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId
	//
	//userDB[user.Id] = user
	return nil
}
