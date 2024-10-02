package handler

import (
	"net/http"

	"svenvermeulen/platform-go-challenge/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FavouritesHandler struct {
}

func NewFavouritesHandler() *FavouritesHandler {
	return &FavouritesHandler{}
}

// GetUserFavourites godoc
//
// @Summary     Get user favourites
// @Description gets list of favourite
// @Tags        favourites
// @Produce     json
// @Success     200	{object}	[]model.Favourite
// @Failure     400
// @Failure     404
// @Failure     500
// @Router      /favourites [get]
func (h FavouritesHandler) GetFavourites(c *gin.Context) {
	_  = c.Param("userid")

	result := model.UserFavourites{
		Charts: []model.Chart{
			{
				Id: uuid.New(),
				Title: "My first favourite chart",
				XAxisTitle: "Time",
				YAxisTitle: "Number of clicks",
				DataPoints: []model.DataPoint{
					{X: 0.0, Y: 10.0},
					{X: 1.0, Y: 12.0},
					{X: 2.0, Y: 15.0},
				},
			},
		},
		Insights: []model.Insight{
			{
				Id: uuid.New(),
				Description: "40% of millenials spend more than 3 hours on social media daily",
			},
			{
				Id: uuid.New(),
				Description: "20% of boomers make online purchases at least once a week",
			},
		},
		Audiences: []model.Audience{
			{
				Id: uuid.New(),
				Gender: 'm',
				BirthCountry: "Netherlands",
				AgeFrom: 45,
				AgeTo: 55,
				HoursSpentFrom: 4,
				HoursSpentTo: 8,
				PurchasesFrom: 0,
				PurchasesTo: 0,
			},
		},
	}
	c.IndentedJSON(http.StatusOK, result)
}
