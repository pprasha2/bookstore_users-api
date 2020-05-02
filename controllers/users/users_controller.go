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

//dto - data transfer object

func getUserID(userIDParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIDParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("invalid userId, userId should be a number")

	}
	return userId, nil
}

//GetUsers fetch the user based on user_id
func Get(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		//TODO Handle User creation error

		c.JSON(getErr.Status, getErr)
		return

	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
	//c.String(http.StatusNotImplemented, "Implement me!")
}

//CreateUser create a new user
//It fetches the parameters to process the request and send it to correspondig service
func Create(c *gin.Context) {
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

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))

}

//SearchUser searches a user
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO return Bad request
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		//TODO Handle User creation error

		c.JSON(err.Status, err)
		return

	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func Delete(c *gin.Context) {
	userId, idErr := getUserID(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
