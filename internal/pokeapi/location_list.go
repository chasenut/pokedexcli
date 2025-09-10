package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (locations RespLocationPage, err error) {
	url := base_url + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var locationResp RespLocationPage
	// chech if cached
	if cached, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(cached, &locationResp)
		if err != nil {
			return RespLocationPage{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationPage{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationPage{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationPage{}, err
	}
	c.cache.Add(url, data) // add to cache

	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespLocationPage{}, err
	}
	
	return locationResp, nil
}
