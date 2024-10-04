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


	


	// set up router to map http routes to handler functions
	router := SetupRouter()
	router.Run("localhost:8086")
}

func SetupRouter() *gin.Engine {
	// set up repositories
	favouriteRepo := favourite.NewRepository()
	audienceRepo := audience.NewRepository()
	chartRepo := chart.NewRepository()
	insightRepo := insight.NewRepository()



	// TODO: FOR QUICK AND DIRTY TEST. REMOVE.
	userId, _ := uuid.Parse("609dac9c-ac79-4dc8-a1f5-f2af7a5519cf")
	favouriteRepo.AddFavourite(userId, "audience I like 1", uuid.New(), "audience")
	favouriteRepo.AddFavourite(userId, "chart I like 1", uuid.New(), "chart")
	favouriteRepo.AddFavourite(userId, "insight I like 1", uuid.New(), "insight")
	favouriteRepo.AddFavourite(userId, "audience I like 2", uuid.New(), "audience")
	favouriteRepo.AddFavourite(userId, "audience I like 3", uuid.New(), "audience")
	favouriteRepo.AddFavourite(userId, "chart I like 2", uuid.New(), "chart")
	favouriteRepo.AddFavourite(userId, "chart I like 3", uuid.New(), "chart")
	favouriteRepo.AddFavourite(userId, "insight I like 2", uuid.New(), "insight")
	favouriteRepo.AddFavourite(userId, "insight I like 3",  uuid.New(), "insight")


	// set up http handlers
	favouritesHandler := handler.NewFavouritesHandler(favouriteRepo, audienceRepo, chartRepo, insightRepo)


	log.Info("Setting up gin router")
	router := gin.Default()

	// setup routes
	router.GET("/favourites", favouritesHandler.GetFavourites)

	return router
}
