package favourite

import (
	"github.com/google/uuid"
)

type FavouriteEntry struct {
	FavouriteId  uuid.UUID
	ResourceType string
}

type FavouriteEntries []FavouriteEntry

type Repository struct {
	// keys in the map are the uniqe ids of the users
	favourites map[uuid.UUID]FavouriteEntries
}

func NewRepository() *Repository {
	return &Repository{
		favourites: make(map[uuid.UUID]FavouriteEntries, 100),
	}
}

func (r *Repository) GetFavourites(userid uuid.UUID, offset int, pageSize int) FavouriteEntries {
	if result, ok := r.favourites[userid]; !ok {
		return make(FavouriteEntries, 0, 100)
	} else {
		if offset > len(result) {
			return FavouriteEntries{}
		}

		high := offset + pageSize
		if high > len(result) {
			high = len(result)
		}
		return result[offset:high]
	}
}

func (r *Repository) AddFavourite(userId uuid.UUID, favouriteId uuid.UUID, favouriteType string) {
	if _, ok := r.favourites[userId]; !ok {
		r.favourites[userId] = make(FavouriteEntries, 0, 100)
	}

	// not checking for duplicate entries
	r.favourites[userId] = append(r.favourites[userId], FavouriteEntry{
		FavouriteId:  favouriteId,
		ResourceType: favouriteType,
	})
}
