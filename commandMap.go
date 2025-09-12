package main

import (
	"errors"
	"fmt"
)

func commandMapf(c *config, args []string) error {
	locationResp, err := c.pokeapiClient.GetLocations(c.nextLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locationResp.Next
	c.prevLocationURL = locationResp.Previous
	
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config, args []string) error {
	if c.prevLocationURL == nil {
		return errors.New("You are on the first page")
	}

	locationResp, err := c.pokeapiClient.GetLocations(c.prevLocationURL)
	if err != nil {
		return err
	}

	c.nextLocationURL = locationResp.Next
	c.prevLocationURL = locationResp.Previous
	
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

