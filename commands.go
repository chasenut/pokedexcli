package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name 		string
	description string
	callback 	func() error
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	helpMessage := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`
	fmt.Println(helpMessage)
	return nil
}

var commands = map[string]cliCommand{
	"exit": cliCommand{
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
	"help": cliCommand{
		name: "help",
		description: "Print the help dialog",
		callback: commandHelp,
	},
}
