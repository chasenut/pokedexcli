package main

import (
	"errors"
	"fmt"
)

func commandInspect(c *config, args []string) error {
	if len(args) == 0 {
		return errors.New("Please provide a Pokémon name")
	}
	name := args[0]
	pokemon, ok := c.pokedex[name]
	if !ok {
		return errors.New("You have not caught that Pokémon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
