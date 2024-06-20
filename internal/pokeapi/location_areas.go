package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Client) ListLocationAreas() (locationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

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

	locationAreas := locationAreasResp{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return locationAreasResp{}, err
	}

	return locationAreas, nil
}
