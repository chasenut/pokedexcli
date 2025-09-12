package pokeapi

import (
	"net/http"
	"time"

	"github.com/chasenut/pokedexcli/internal/pokecache"
)

type Client struct {
	cache 		pokecache.Cache
	httpClient 	http.Client
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		cache: *cache,
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
