package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mathiaskluge/pokedex-cli/internal/pokecache"
)

type locationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(paginationURL *string, cache *pokecache.Cache) (locationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if paginationURL != nil {
		fullURL = *paginationURL
	}

	if body, ok := cache.Get(fullURL); ok {
		locationAreas := locationAreasResp{}
		err := json.Unmarshal(body, &locationAreas)
		if err != nil {
			return locationAreasResp{}, err
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreasResp{}, err
	}
	cache.Add(fullURL, body)

	locationAreas := locationAreasResp{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return locationAreasResp{}, err
	}
	cache.Add(fullURL, body)
	return locationAreas, nil

}
