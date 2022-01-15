package main

import (
	"fmt"
	"pokedex-go/pokemon"
)

func main() {
	//pokemon to get
	pokeToGet := 151

	//set concurrency limit
	sem := pokemon.NewSem(20)

	pokedex := pokemon.BuildPokedex(sem, pokeToGet)

	for _, poke := range pokedex {
		fmt.Println(poke.String())
	}

}
