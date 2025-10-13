package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaList,error) {
	fullURL := BaseURL + "/location-area"
	locations := LocationAreaList{}
	// if we get a pageUrl make it the new url
	if pageURL != nil {
		fullURL = *pageURL
	}
	// try the cache
	val, ok := c.cache.Get(fullURL)
	if ok {
		if err := json.Unmarshal(val, &locations); err != nil{
			return LocationAreaList{}, err
		}
		return locations, nil
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
	//cache the bytes for future fast readings
	c.cache.Add(fullURL, data)

	// make empty struct and unmarashal the []bytes into the struct
	if err = json.Unmarshal(data, &locations); err != nil {
		return LocationAreaList{}, err
	}
	
	return locations, nil
}
