package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pprasha2/bookstore_users-api/logger"
)

var (
	router = gin.Default()
)

//StartApplication start App
func StartApplication() {
	mapUrls()
	logger.Info("about to start the application ...")
	router.Run(":8000")
}
