package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		errors.New("Please provide a Pokémon name")
	}
	name := args[0]
	pokemon, err := c.pokeapiClient.GetPokemon(name)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not find Pokémon named %s\n", pokemon.Name))
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	
	res := rand.Intn(pokemon.BaseExperience)
	if res < 50 {
		c.PokedexAdd(pokemon.Name, pokemon)
		fmt.Printf("%s was caught!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}
