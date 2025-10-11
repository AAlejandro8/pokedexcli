package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
)

const baseURL = "https://pokeapi.co/api/v2/"

func getLocationAreas(url string) (LocationAreaList,error) {
	// make the request
	if url == "" {
		url = baseURL + "location-area?limit=20"
	}

	res, err := http.NewRequest("GET",url, nil)
	if err != nil {
		return LocationAreaList{}, err
	}
	defer res.Body.Close()

	// make the struct and populate it
	var locations LocationAreaList
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaList{}, err
	}

	if err = json.Unmarshal(body, &locations) ;err != nil {
		return LocationAreaList{}, err
	}

	return locations, err
}

