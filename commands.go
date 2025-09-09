package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/chasenut/pokedexcli/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationURL 	*string
	prevLocationURL 	*string
}

type cliCommand struct {
	name 		string
	description string
	callback 	func(c *config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	helpMessage := `
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`
	fmt.Println(helpMessage)
	return nil
}

func commandMapf(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous
	
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you are on the first page")
	}

	locationResp, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous
	
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

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
	"map": cliCommand{
		name: "map",
		description: "Get the next page of locations",
		callback: 
		commandMapf,
	},
	"mapb": cliCommand{
		name: "map",
		description: "Get the previous page of locations",
		callback: commandMapb,
	},
}
