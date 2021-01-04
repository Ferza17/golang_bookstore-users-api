package app

import (
	"github.com/gin-gonic/gin"
	Log "github.com/ferza17/golang_bookstore-users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	Log.Info("The Application about to Start...")
	_ = router.Run(":8081")
}

