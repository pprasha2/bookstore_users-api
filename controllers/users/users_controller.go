package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pprasha2/bookstore_users-api/domain/users"
	"github.com/pprasha2/bookstore_users-api/services"
	"github.com/pprasha2/bookstore_users-api/utils/errors"
)

//GetUsers fetch the user based on user_id
func GetUsers(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid userId, userId should be a number")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		//TODO Handle User creation error

		c.JSON(getErr.Status, getErr)
		return

	}
	c.JSON(http.StatusOK, user)
	//c.String(http.StatusNotImplemented, "Implement me!")
}

//CreateUser create a new user
//It fetches the parameters to process the request and send it to correspondig service
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO return Bad request
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO Handle User creation error

		c.JSON(saveErr.Status, saveErr)
		return

	}

	fmt.Println(user)

	c.JSON(http.StatusCreated, result)

}

//SearchUser searches a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
