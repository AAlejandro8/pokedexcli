package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaList,error) {
	fullURL := BaseURL + "/location-area"
	// if we get a pageUrl make it the new url
	if pageURL != nil {
		fullURL = *pageURL
	}
	// make the request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaList{}, err
	}
	// call the api with the client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaList{}, err
	}
	// close the body
	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return LocationAreaList{}, fmt.Errorf("error status code: %v ", resp.StatusCode)
	}
	// get the data into []bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaList{}, err
	}
	// make empty struct and unmarashal the []bytes into the struct
	locations := LocationAreaList{}
	if err = json.Unmarshal(data, &locations); err != nil {
		return LocationAreaList{}, nil
	}
	
	return locations, nil
}
