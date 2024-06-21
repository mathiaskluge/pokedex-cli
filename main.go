package main

import (
	"time"

	"github.com/mathiaskluge/pokedex-cli/internal/pokeapi"
	"github.com/mathiaskluge/pokedex-cli/internal/pokecache"
)

type config struct {
	pokeapiClient       pokeapi.Client
	pokeCache           pokecache.Cache
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		pokeCache:     pokecache.NewCache(time.Minute * 5),
	}

	startRepl(&cfg)
}
