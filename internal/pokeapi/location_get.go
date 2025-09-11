package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location *string) (pokemons Location, err error) {
	url := base_url + "/location-area/" + *location

	var locationPokemons Location
	if cached, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(cached, &locationPokemons)
		if err != nil {
			return Location{}, err
		}
		return locationPokemons, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationPokemons)
	if err != nil {
		return Location{}, err
	}
	return locationPokemons, nil
}
