package users

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/pprasha2/bookstore_users-api/domain/users"
	"github.com/pprasha2/bookstore_users-api/services"
)

//GetUsers fetch the user based on user_id
func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

//CreateUser create a new user
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO handle error
		fmt.Println("in read error! bad json")
		fmt.Println(err.Error())
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO handle json error
		fmt.Println("in unmarshal json,incorrect payload")
		fmt.Println(err.Error())
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Error while creating a user
		//TODO Handle
		return

	}
	fmt.Println(result)
	fmt.Println(user)
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.String(http.StatusNotImplemented, "Implement me!")

}

//SearchUser searches a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
