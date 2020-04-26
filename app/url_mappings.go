package app

import (
	"github.com/pprasha2/bookstore_users-api/controllers/ping"
	"github.com/pprasha2/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUsers)
	router.GET("/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
