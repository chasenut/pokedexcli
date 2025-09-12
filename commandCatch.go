package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please provide a Pokémon name")
	}
	name := args[0]
	pokemon, err := c.pokeapiClient.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("Could not find Pokémon named %s\n", pokemon.Name)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	
	res := rand.Intn(pokemon.BaseExperience)
	if res < 50 {
		c.PokedexAdd(pokemon.Name, pokemon)
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
		return nil
	}
	fmt.Printf("%s escaped!\n", pokemon.Name)
	return nil
}
