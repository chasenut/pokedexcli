package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (locations RespLocationsPage, err error) {
	url := base_url + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationResp RespLocationsPage
	// chech if cached
	if cached, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(cached, &locationResp)
		if err != nil {
			return RespLocationsPage{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsPage{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsPage{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationsPage{}, err
	}
	c.cache.Add(url, data)

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespLocationsPage{}, err
	}
	
	return locationResp, nil
}
