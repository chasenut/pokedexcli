package main

import (
	"time"

	"github.com/chasenut/pokedexcli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second * 5)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startREPL(cfg)
}
