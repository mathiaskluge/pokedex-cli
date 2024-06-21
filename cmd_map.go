package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL, &cfg.pokeCache)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL, &cfg.pokeCache)
	if err != nil {
		return err
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous

	return nil
}
