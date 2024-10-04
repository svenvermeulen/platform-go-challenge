package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"svenvermeulen/platform-go-challenge/internal/handler"
	"svenvermeulen/platform-go-challenge/internal/repository/audience"
	"svenvermeulen/platform-go-challenge/internal/repository/chart"
	"svenvermeulen/platform-go-challenge/internal/repository/favourite"
	"svenvermeulen/platform-go-challenge/internal/repository/insight"
)

// @title Favourites API
func main() {
	// set up repositories
	favouriteRepo := favourite.NewRepository()
	audienceRepo := audience.NewRepository()
	chartRepo := chart.NewRepository()
	insightRepo := insight.NewRepository()

	// TODO: FOR QUICK AND DIRTY TEST. REMOVE.
	userId, _ := uuid.ParseBytes([]byte{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7})
	favouriteRepo.AddFavourite(userId, uuid.New(), "audience")
	favouriteRepo.AddFavourite(userId, uuid.New(), "chart")
	favouriteRepo.AddFavourite(userId, uuid.New(), "insight")
	favouriteRepo.AddFavourite(userId, uuid.New(), "audience")
	favouriteRepo.AddFavourite(userId, uuid.New(), "audience")
	favouriteRepo.AddFavourite(userId, uuid.New(), "chart")
	favouriteRepo.AddFavourite(userId, uuid.New(), "chart")
	favouriteRepo.AddFavourite(userId, uuid.New(), "insight")
	favouriteRepo.AddFavourite(userId, uuid.New(), "insight")

	// set up http handlers
	favouritesHandler := handler.NewFavouritesHandler(favouriteRepo, audienceRepo, chartRepo, insightRepo)

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
