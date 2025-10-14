package pokeapi

import (
	"net/http"
	"time"
	pokeapi "github.com/AAlejandro8/pokedexcli/internal/pokeapi/pokecache"
)

type Client struct {
	httpClient http.Client
	cache *pokeapi.Cache
}

func NewClient() Client{
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokeapi.NewCache(time.Second * 5),
	}
}