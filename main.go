package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

//pokemon struct representation
type pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Wieght int    `json:"weight"`

	Type []struct {
		ElemType struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`

	AttackMove []struct {
		Moves struct {
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
}

//stringify
func (p *pokemon) String() string {
	return fmt.Sprintf(
		"ID: %d - NAME: %s - WEIGHT: %d - TYPE: %s - ATTACK-MOVE: %s",
		p.ID, p.Name, p.Wieght, p.Type[0].ElemType.Name, p.AttackMove[0].Moves.Name)
}

//getPokemon performs a GET request to the
//pokemon api, marshals the needed data and
//returns a correlating pokemon
func getPokemon(uri string) pokemon {
	var p pokemon
	fmt.Printf("ASYNC Request: %s\n", uri)
	resp, err := http.Get(uri)
	if err != nil {
		log.Println(err)
		return pokemon{}
	}

	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		return pokemon{}
	}

	return p
}

func main() {
	wg := sync.WaitGroup{}
	pokedex := make(chan pokemon, 151)

	for i := 1; i <= 151; i++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			pokedex <- getPokemon(
				fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", p),
			)
		}(i)
	}

	wg.Wait()
	close(pokedex)

	for poke := range pokedex {
		fmt.Println(poke.String())
	}

}
