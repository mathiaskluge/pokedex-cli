package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func callbackMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		fmt.Println("Problem fetching data:", err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Problem reading the received data:", err)
		return err
	}

	var loc locations
	if err := json.Unmarshal(body, &loc); err != nil {
		fmt.Println("Problem decoding response:", err)
		return err
	}

	fmt.Printf("Number of locations: %d\n", loc.Count)
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}

	return nil
}
