package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache *Cache
}

func NewClient() Client{
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: NewCache(time.Second * 5),
	}
}