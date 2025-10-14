package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaList, error) {
	fullURL := BaseURL + "/location-area"
	// if we get a pageUrl make it the new url
	if pageURL != nil {
		fullURL = *pageURL
	}
	// try the cache
	val, ok := c.cache.Get(fullURL)
	if ok {
		locations := LocationAreaList{}
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
	locations := LocationAreaList{}
	if err = json.Unmarshal(data, &locations); err != nil {
		return LocationAreaList{}, err
	}
	
	return locations, nil
}

func (c *Client) ExploreLocation(location string) (ExploredArea, error) {
	fullURL := BaseURL + "/location-area/" + location

	// try the cache first 
	val, ok := c.cache.Get(fullURL) 
		if ok {
			// cache hit
			encounters := ExploredArea{}
			if err := json.Unmarshal(val, &encounters); err != nil {
				return ExploredArea{}, err
			}
			return encounters, nil
		}

	// make the request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return ExploredArea{}, err
	}

	// make the call 
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploredArea{}, err
	}
	// close body at function end
	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return ExploredArea{}, fmt.Errorf("error status code: %v ", resp.StatusCode)
	}

	// get the data into []bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploredArea{}, err
	}
	// cache findings for fast rerevials 
	c.cache.Add(fullURL,data)

	// make empty encounters and unmarshal into it
	encounters := ExploredArea{}
	if err = json.Unmarshal(data, &encounters); err != nil {
		return ExploredArea{}, err
	}

	return encounters, nil
} 
