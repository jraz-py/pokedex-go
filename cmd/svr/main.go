package main

import (
	"fmt"
	"pokedex-go/pokemon"
)

const pokeUrl = "https://pokeapi.co/api/v2/pokemon"

func main() {
	//pokemon to get
	pokeToGet := 151

	//set concurrency limit
	sem := pokemon.NewSem(100)

	pokedex := pokemon.BuildPokedex(sem, pokeToGet, pokeUrl)

	for _, poke := range pokedex {
		fmt.Println(poke.String())
	}

}
