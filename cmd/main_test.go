package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHappyPath(t* testing.T) {
	
	router := SetupRouter()

	w := httptest.NewRecorder()



	// GIVEN a user with a few favourite items
	req, err := http.NewRequest("POST", "/favourites", nil)
	require.NoError(t, err, "error creating http request")

	router.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
	

	require.Equal(t, 1, len(result))

	// WHEN I retrieve the user's favourites
	result := []model.ItemPrice{}
	if err := json.NewDecoder(w.Result().Body).Decode(&result); err != nil {
		t.Fatalf("could not decode response body: %s", err)
	}
	// THEN I get a list of audiences, insights and charts
	
}