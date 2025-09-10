package main

import (
	"time"

	"github.com/chasenut/pokedexcli/internal/pokeapi"
	"github.com/chasenut/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(time.Second * 5, &cache)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startREPL(cfg)
}
