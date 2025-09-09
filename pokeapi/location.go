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
	var locationResp RespLocationPage
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespLocationPage{}, err
	}
	
	return locationResp, nil
}
