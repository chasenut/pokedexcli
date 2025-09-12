package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a Pokémon name")
		return nil
	}
	name := args[0]
	pokemon, err := c.pokeapiClient.GetPokemon(&name)
	if err != nil {
		fmt.Printf("Could not find Pokémon named %s\n", name)
		return nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	
	base_exp := min(pokemon.BaseExperience, 300) 
	rng := rand.Intn(400)
	if base_exp < rng {
		c.PokedexAdd(name, pokemon)
		fmt.Printf("%s was caught!\n", name)
		return nil
	}
	fmt.Printf("%s escaped!\n", name)
	return nil
}
