package main

import (
	"fmt"
)

func commandMapf(cfg *config, args []string) error {
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

func commandMapb(cfg *config, args []string) error {
	if cfg.prevLocationURL == nil {
		fmt.Println("You are on the first page")
		return nil
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

