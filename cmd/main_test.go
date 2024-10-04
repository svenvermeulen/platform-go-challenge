package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"svenvermeulen/platform-go-challenge/pkg/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHappyPath(t* testing.T) {
	router := SetupRouter()

	// GIVEN a user with a favourite item
	f := model.UserFavouriteShort{
		Description:  "favourite description",
		ResourceType: "chart",
		Id:           uuid.New(),
	}

	bodyBytes, err := json.Marshal(f)
    if err != nil {
        log.Fatal(err)
    }

	body := bytes.NewBuffer(bodyBytes)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/favourites", body)
	req.Header.Add("Authorization", getAuthToken())
	require.NoError(t, err, "error creating http request")
	router.ServeHTTP(w, req)
	require.Equal(t, 201, w.Code)

	// WHEN I retrieve the user's favourites
	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/favourites", nil)
	req.Header.Add("Authorization", getAuthToken())
	require.NoError(t, err, "error creating http request")
	router.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	result := []model.UserFavourite{}
	if err := json.NewDecoder(w.Result().Body).Decode(&result); err != nil {
		t.Fatalf("could not decode response body: %s", err)
	}

	// THEN I get a list of audiences, insights and charts
	assert.Equal(t, 1, len(result))
}

func TestDeletion(t* testing.T) {
	router := SetupRouter()

	var uuidToDelete uuid.UUID

	// GIVEN a user with a few favourite items
	for i := range 2 {
		f := model.UserFavouriteShort{
			Description:  fmt.Sprintf("favourite %d description", i),
			ResourceType: "chart",
			Id:           uuid.New(),
		}
		uuidToDelete = f.Id

		bodyBytes, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
		}

		body := bytes.NewBuffer(bodyBytes)
	
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/favourites", body)
		req.Header.Add("Authorization", getAuthToken())
		require.NoError(t, err, "error creating http request")
		router.ServeHTTP(w, req)
		require.Equal(t, 201, w.Code)
	}
	// WHEN I delete one favourite
	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/favourites/%v", uuidToDelete), nil)
	req.Header.Add("Authorization", getAuthToken())
	require.NoError(t, err, "error creating http request")
	router.ServeHTTP(w, req)
	require.Equal(t, 204, w.Code)

	// WHEN I retrieve the user's favourites
	w = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/favourites", nil)
	req.Header.Add("Authorization", getAuthToken())
	require.NoError(t, err, "error creating http request")
	router.ServeHTTP(w, req)
	require.Equal(t, 200, w.Code)

	result := []model.UserFavourite{}
	if err := json.NewDecoder(w.Result().Body).Decode(&result); err != nil {
		t.Fatalf("could not decode response body: %s", err)
	}

	// THEN I get a list of audiences, insights and charts
	assert.Equal(t, 1, len(result))
}

func getAuthToken() string {
	return "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJmYXZvdXJpdGVzIiwidXNlcmlkIjoiNjA5ZGFjOWMtYWM3OS00ZGM4LWExZjUtZjJhZjdhNTUxOWNmIiwiaWF0IjoxNzI4MDM2ODYyLCJleHAiOjE3MjgwNTA0NjJ9.B23I_5l52uMC77Ueqq8KGeRickJ6Vy_iqKrBXWBKy30"
}