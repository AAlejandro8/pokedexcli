package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client)GetPokemonInfo(pokemon string) (Pokemon, error) {
	fullURL := BaseURL + "/pokemon/" + pokemon
	// make the request

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// make the client call the api
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	} 
	defer resp.Body.Close()

	// read the resp into []bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// make struct to populate
	pokemonToReturn := Pokemon{}
	if err = json.Unmarshal(data, &pokemonToReturn); err != nil {
		return Pokemon{}, err
	}

	return pokemonToReturn, nil

}