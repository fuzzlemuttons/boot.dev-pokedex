package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(pokemon string) (Pokemon, error) {
	url := pokeApiBaseUrl + "/pokemon/" + pokemon

	locationsResp := Pokemon{}
	cacheVal, exists := c.cache.Get(url)
	if exists {
		err := json.Unmarshal(cacheVal, &locationsResp)
		if err != nil {
			return Pokemon{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, dat)

	return locationsResp, nil
}
