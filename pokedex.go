package main

import "github.com/chasenut/pokedexcli/internal/pokeapi"

func (cfg *config) PokedexAdd(name string, pokemon pokeapi.Pokemon) {
	cfg.pokedex[name] = pokemon
}

func (cfg *config) PokedexList() []pokeapi.Pokemon {
	pokemons := []pokeapi.Pokemon{}
	for _, p := range cfg.pokedex {
		pokemons = append(pokemons, p)
	}
	return pokemons
}

