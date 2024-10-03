package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"svenvermeulen/platform-go-challenge/api/handler"
)

// @title Favourites API
func main() {
	router := SetupRouter()
	router.Run("localhost:8086")
}

func SetupRouter() *gin.Engine {
	log.Info("Setting up gin router")
	router := gin.Default()
	favouritesHandler := handler.NewFavouritesHandler()

	// setup routes
	router.GET("/favourites", favouritesHandler.GetFavourites)

	return router
}
