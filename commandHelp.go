package main

import "fmt"

func commandHelp(c *config, args []string) error {
	helpMessage := `
Welcome to the Pokedex!
Usage:
`
	fmt.Println(helpMessage)
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	fmt.Println()

	return nil
}

