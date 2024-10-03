package handler

import (
	"net/http"

	"svenvermeulen/platform-go-challenge/internal/repository/audience"
	"svenvermeulen/platform-go-challenge/internal/repository/chart"
	"svenvermeulen/platform-go-challenge/internal/repository/insight"
	"svenvermeulen/platform-go-challenge/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FavouritesHandler struct {
	audienceRepository *audience.Repository
	chartRepository    *chart.Repository
	insightRepository  *insight.Repository
}

func NewFavouritesHandler(audienceRepository *audience.Repository, chartRepository *chart.Repository, insightRepository *insight.Repository) *FavouritesHandler {
	return &FavouritesHandler{
		audienceRepository: audienceRepository,
		chartRepository:    chartRepository,
		insightRepository:  insightRepository,
	}
}

// GetUserFavourites godoc
//
// @Summary     Get user favourites
// @Description gets list of favourites for a specified user
// @Tags        favourites
// @Produce     json
// @Success     200	{object}	[]model.UserFavourite
// @Failure     400
// @Failure     404
// @Failure     500
// @Router      /favourites/:userid [get]
func (h FavouritesHandler) GetFavourites(c *gin.Context) {
	_ = c.Param("userid") // TODO: Should come from session / token, not request

	// What I'll probably do is
	// - get "start after" and "page size" parameters from the request, for paging
	// - get the next <page_size> favourites for the user from the db
	//   - these will look like <uuid> <resource type>
	// - split the favourites into 3 lists of ids, one per resource type
	// - retrieve the resources by ID from 3 repos
	//   - let's say insights and audiences are simply retrieved from a RDBMS
	//   - let's say charts are generated by an external service and take a bit longer to return
	// - I wait for these 3 calls to finish, then stitch together the results and return them in JSON format

	// This will show
	// - using concurrency to retrieve data from multiple sources
	// - using result paging for more responsive list endpoints
	// - using userid from session/token for security
	// Maybe:
	// - graceful degradation. If the call to charts service asks for too many charts, it can time out.
	//   the call to "favourites" could still work though, with the "charts" being empty.

	// TODO: These ids should come from some repository which reads "starred items"
	insightIDs := []uuid.UUID{uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New()}
	insights := h.insightRepository.GetInsights(insightIDs)

	audienceIDs := []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}
	audiences := h.audienceRepository.GetAudiences(audienceIDs)

	chartIDs := []uuid.UUID{uuid.New(), uuid.New()}
	charts := h.chartRepository.GetCharts(chartIDs)

	result := model.UserFavourites{
		Audiences: audiences,
		Charts:    charts,
		Insights:  insights,
	}
	c.IndentedJSON(http.StatusOK, result)
}
