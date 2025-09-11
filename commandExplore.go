package main

import (
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please provide the name of the location")
		return nil
	}
	locationName := args[0]
	locationPokemons, err := c.pokeapiClient.GetLocation(&locationName)
	if err != nil {
		fmt.Printf("Non-existent location \"%v\"\n", locationName)
		return nil
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationName)
	for _, p := range locationPokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}
