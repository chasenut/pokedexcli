package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please provide the name of the location")
	}
	locationName := args[0]
	locationPokemons, err := c.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return fmt.Errorf("Non-existent location \"%v\"\n", locationName)
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationName)
	for _, p := range locationPokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", p.Pokemon.Name)
	}
	return nil
}
