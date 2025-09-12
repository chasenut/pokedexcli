package main

import "fmt"

func commandPokedex(c *config, args []string) error {
	pokemons := c.PokedexList()
	if len(pokemons) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, p := range pokemons {
		fmt.Printf("  - %s\n", p.Name)
	}
	return nil
}
