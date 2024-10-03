package favourite

import (
	"github.com/google/uuid"
)

type Repository struct {
	// keys in the map are the uniqe ids of the users.
	// values in outer map are maps which map resource id to resource type (audience|chart|insight)
	favourites map[uuid.UUID]map[uuid.UUID]string
}

func NewRepository() *Repository {
	return &Repository{
		favourites: make(map[uuid.UUID]map[uuid.UUID]string, 100),
	}
}

func (r *Repository) GetFavourites(userid uuid.UUID) map[uuid.UUID]string {
	if result, ok := r.favourites[userid]; !ok {
		return make(map[uuid.UUID]string)
	} else {
		return result
	}
}

func (r *Repository) AddFavourite(userId uuid.UUID, favouriteId uuid.UUID, favouriteType string) {
	if _, ok := r.favourites[userId]; !ok {
		r.favourites[userId] = make(map[uuid.UUID]string, 100)
	}
	r.favourites[userId][favouriteId] = favouriteType
}
