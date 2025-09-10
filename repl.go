package main

import (
	//	"fmt"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chasenut/pokedexcli/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationURL 	*string
	prevLocationURL 	*string
}

func startREPL(c *config) {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for true {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("error, scanner run out of tokens or raised an error (i guess)")
			continue
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		if cmd, ok := getCommands()[words[0]]; ok {
			err := cmd.callback(c)
			if err != nil {
				fmt.Printf("Error during function callback: %s\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	splitted := strings.Fields(lowered)
	return splitted
}

type cliCommand struct {
	name 		string
	description string
	callback 	func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapb,
		},
	}
}
