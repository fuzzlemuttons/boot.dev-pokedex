package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (RespLocation, error) {
	url := pokeApiBaseUrl + "/location-area/" + location

	locationsResp := RespLocation{}
	cacheVal, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(cacheVal, &locationsResp)
		if err != nil {
			return RespLocation{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocation{}, err
	}
	c.cache.Add(url, dat)

	return locationsResp, nil
}
