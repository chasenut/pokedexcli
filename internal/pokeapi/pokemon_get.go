package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (pokemon Pokemon, err error) {
	url := base_url + "/pokemon/" + name

	var respPokemon Pokemon
	if cached, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(cached, &respPokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return respPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &respPokemon)	
	if err != nil {
		return Pokemon{}, err
	}

	return respPokemon, nil
}
