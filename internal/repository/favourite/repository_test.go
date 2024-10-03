package favourite

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetNonexistingUser(t *testing.T) {
	// if user is not in the data set, empty map[uuid]string should be returned

	// GIVEN an empty repo
	r := NewRepository()

	// WHEN I retrieve a user's favourites
	f := r.GetFavourites(uuid.New())

	// THEN I get an empty map
	assert.NotNil(t, f, "GetFavourites must return empty map for non-existent user uuid")
}

func TestAddAndRetrieveCorrectFavourite(t *testing.T) {
	// if I add a favourite for two users, the favourite for the correct user should be returned
	user1Id := uuid.New()
	favourite1Id := uuid.New()

	user2Id := uuid.New()
	favourite2Id := uuid.New()

	// GIVEN a repo with two users' favourites in it
	r := NewRepository()
	r.AddFavourite(user1Id, favourite1Id, "chart")
	r.AddFavourite(user2Id, favourite2Id, "insight")

	// WHEN I retrieve each user's favourites
	f1 := r.GetFavourites(user1Id)
	f2 := r.GetFavourites(user2Id)

	// THEN I get the correct favourites for each user
	assert.Equal(t, f1, map[uuid.UUID]string {
		favourite1Id: "chart",
	} )

	assert.Equal(t, f2, map[uuid.UUID]string {
		favourite2Id: "insight",
	} )
}