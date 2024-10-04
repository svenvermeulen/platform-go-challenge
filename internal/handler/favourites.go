package handler

import (
	"net/http"
	"strconv"
	"sync"

	"svenvermeulen/platform-go-challenge/internal/repository/audience"
	"svenvermeulen/platform-go-challenge/internal/repository/chart"
	"svenvermeulen/platform-go-challenge/internal/repository/favourite"
	"svenvermeulen/platform-go-challenge/internal/repository/insight"
	"svenvermeulen/platform-go-challenge/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FavouritesHandler struct {
	favouriteRepository *favourite.Repository
	audienceRepository  *audience.Repository
	chartRepository     *chart.Repository
	insightRepository   *insight.Repository
}

func NewFavouritesHandler(favouriteRepository *favourite.Repository,
	audienceRepository *audience.Repository,
	chartRepository *chart.Repository,
	insightRepository *insight.Repository) *FavouritesHandler {
	return &FavouritesHandler{
		favouriteRepository: favouriteRepository,
		audienceRepository:  audienceRepository,
		chartRepository:     chartRepository,
		insightRepository:   insightRepository,
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
func (h *FavouritesHandler) GetFavourites(c *gin.Context) {
	userId, _ := uuid.ParseBytes([]byte{0, 1, 2, 3, 4, 5, 6, 7, 0, 1, 2, 3, 4, 5, 6, 7}) // TODO: Should come from session / token, not request

	var err error
	var offset int
	var pageSize int

	if offset, err = strconv.Atoi(c.Query("offset")); err!=nil {
		offset = 0
	}
	if pageSize, err = strconv.Atoi(c.Query("pagesize")); err!=nil {
		pageSize = 10
	}

	// TODO
	// - some logging
	// - auth
	// - move quick and dirty test to nice automated test
	// - README
	// - swagger stuff

	// shown:
	// - using concurrency to retrieve data from multiple sources
	// - using result paging for more responsive list endpoints
	// - using userid from session/token for security
	// Maybe:
	// - graceful degradation. If the call to charts service asks for too many charts, it can time out.
	//   the call to "favourites" could still work though, with the "charts" being empty.
	userFavourites := h.favouriteRepository.GetFavourites(userId, offset, pageSize)

	// Get favourite items for current user and extract these into 3 slices of IDs 
	// IDs are then used to query the various repositories
	// Then I wait for all the queries to return and stitch the results back together
	audienceIDs, insightIDs, chartIDs := h.splitUserFavourites(userFavourites)

	var wg sync.WaitGroup
	wg.Add(3)

	// AUDIENCES
	var audiences map[uuid.UUID]*model.Audience

	go func() {
		defer wg.Done()
		audiences = h.audienceRepository.GetAudiences(audienceIDs)
	}()

	// CHARTS
	var charts map[uuid.UUID]*model.Chart

	go func() {
		defer wg.Done()
		charts = h.chartRepository.GetCharts(chartIDs)
	}()

	// INSIGHTS
	var insights map[uuid.UUID]*model.Insight

	go func() {
		defer wg.Done()
		insights = h.insightRepository.GetInsights(insightIDs)
	}()
	
	wg.Wait()

	// Stitch the responses from the various data sources back together in the
	// same order as the page of favourite ids was retrieved from the DB
	result := h.stitchResults(userFavourites, audiences, charts, insights)

	c.IndentedJSON(http.StatusOK, result)
}

func (*FavouritesHandler) stitchResults(userFavourites favourite.FavouriteEntries, 
	audiences map[uuid.UUID]*model.Audience, 
	charts map[uuid.UUID]*model.Chart, 
	insights map[uuid.UUID]*model.Insight) []model.UserFavourite {
	result := make([]model.UserFavourite, 0, len(userFavourites))
	for _, f := range userFavourites {
		switch f.ResourceType {
		case "audience":
			{
				result = append(result, model.UserFavourite{Audience: audiences[f.FavouriteId]})
			}
		case "chart":
			{
				result = append(result, model.UserFavourite{Chart: charts[f.FavouriteId]})
			}
		case "insight":
			{
				result = append(result, model.UserFavourite{Insight: insights[f.FavouriteId]})
			}
		}
	}
	return result
}

// Splits the list of user favourite entries by resource type and returns 3 slices of ids
func (*FavouritesHandler) splitUserFavourites(userFavourites favourite.FavouriteEntries) ([]uuid.UUID, []uuid.UUID, []uuid.UUID) {
	audienceIDs := make([]uuid.UUID, 0, len(userFavourites))
	insightIDs := make([]uuid.UUID, 0, len(userFavourites))
	chartIDs := make([]uuid.UUID, 0, len(userFavourites))

	for _, f := range userFavourites {
		switch f.ResourceType {
		case "audience":
			{
				audienceIDs = append(audienceIDs, f.FavouriteId)
			}
		case "chart":
			{
				chartIDs = append(chartIDs, f.FavouriteId)
			}
		case "insight":
			{
				insightIDs = append(insightIDs, f.FavouriteId)
			}
		default:
			{
				// would log an error here to alert engineer but let request succeed
			}
		}
	}
	return audienceIDs, insightIDs, chartIDs
}
