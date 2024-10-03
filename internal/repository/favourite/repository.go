package favourite

import (
	"github.com/google/uuid"
)

type favouriteEntry struct {
	favouriteId  uuid.UUID
	resourceType string
}

type favouriteEntries []favouriteEntry

type Repository struct {
	// keys in the map are the uniqe ids of the users
	favourites map[uuid.UUID]favouriteEntries
}

func NewRepository() *Repository {
	return &Repository{
		favourites: make(map[uuid.UUID]favouriteEntries, 100),
	}
}

func (r *Repository) GetFavourites(userid uuid.UUID, offset int, pageSize int) favouriteEntries {
	if result, ok := r.favourites[userid]; !ok {
		return make(favouriteEntries, 0, 100)
	} else {
		if offset > len(result) {
			return favouriteEntries{}
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
		r.favourites[userId] = make(favouriteEntries, 0, 100)
	}

	// not checking for duplicate entries
	r.favourites[userId] = append(r.favourites[userId], favouriteEntry{
		favouriteId:  favouriteId,
		resourceType: favouriteType,
	})
}
