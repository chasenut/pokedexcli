package main

import (
	//	"fmt"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	splitted := strings.Fields(lowered)
	return splitted
}

func REPL() {
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
		if cmd, ok := commands[words[0]]; ok == true {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error durin function callback: %s", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		
	}
}

