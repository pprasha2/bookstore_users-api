package users

import (
	"fmt"

	"github.com/pprasha2/bookstore_users-api/datasources/mysql/users_db"
	"github.com/pprasha2/bookstore_users-api/utils/errors"
	mysql_utils "github.com/pprasha2/bookstore_users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail      = "email_UNIQUE"
	queryInsertUser       = "INSERT INTO users(first_name,last_name,email,date_created,status,password) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE id = ?;"
	errorNoRows           = "no rows in result set"
	queryUpdateUser       = "UPDATE users SET first_name=?,last_name=?,email=? WHERE id=?;"
	queryDeleteUser       = "DELETE from users WHERE id=?"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE status=?"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
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
	//user.DateCreated = date_utils.GetNowString()
	insertResult, SaveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if SaveErr != nil {
		fmt.Println(SaveErr)
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

func (user *User) Update() *errors.RestErr {
	fmt.Println("in dao")
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		fmt.Println("in dao error")
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		fmt.Println("in dao execute")
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No user found with status %s", status))
	}
	return results, nil

}
