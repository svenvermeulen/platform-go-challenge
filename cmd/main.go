package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @title Pricing API
func main() {
	router := SetupRouter()
	router.Run("localhost:8086")
}

func SetupRouter() *gin.Engine {
	log.Info("Setting up gin router")
	router := gin.Default()

	// setup routes
	//router.GET("/users",

	return router
}
