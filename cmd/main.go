package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"svenvermeulen/platform-go-challenge/internal/handler"
	"svenvermeulen/platform-go-challenge/internal/repository/audience"
	"svenvermeulen/platform-go-challenge/internal/repository/insight"
)

// @title Favourites API
func main() {
	// set up repositories
	audienceRepo := audience.NewRepository()
	insightRepo := insight.NewRepository()

	// set up http handlers
	favouritesHandler := handler.NewFavouritesHandler(audienceRepo, insightRepo)

	// set up router to map http routes to handler functions
	router := SetupRouter(favouritesHandler)
	router.Run("localhost:8086")
}

func SetupRouter(favouritesHandler *handler.FavouritesHandler) *gin.Engine {
	log.Info("Setting up gin router")
	router := gin.Default()

	// setup routes
	router.GET("/favourites", favouritesHandler.GetFavourites)

	return router
}
