package pokemon

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
)

//pokemon struct representation
type Pokemon struct {
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
func (p *Pokemon) String() string {
	return fmt.Sprintf(
		"ID: %d - NAME: %s - WEIGHT: %d - TYPE: %s - ATTACK-MOVE: %s",
		p.ID, p.Name, p.Wieght, p.Type[0].ElemType.Name, p.AttackMove[0].Moves.Name)
}

//getPokemon performs a GET request to the
//pokemon api, marshals the needed data and
//returns a correlating pokemon
func getPokemon(uri string) Pokemon {
	var p Pokemon
	fmt.Printf("ASYNC Request: %s\n", uri)
	resp, err := http.Get(uri)
	if err != nil {
		log.Println(err)
		return Pokemon{}
	}

	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		return Pokemon{}
	}

	return p
}

func BuildPokedex(sem *sem, pokeToGet int) []Pokemon {
	wg := sync.WaitGroup{}
	pokeChan := make(chan Pokemon, pokeToGet)
	for i := 1; i <= pokeToGet; i++ {
		sem.Aquire()
		wg.Add(1)
		go func(p int) {
			defer sem.Release()
			defer wg.Done()
			pokeChan <- getPokemon(
				fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", p),
			)
		}(i)
	}

	wg.Wait()
	close(pokeChan)

	pokedex := make([]Pokemon, 0)
	for poke := range pokeChan {
		pokedex = append(pokedex, poke)
	}

	sort.SliceStable(pokedex, func(i, j int) bool {
		return pokedex[i].ID < pokedex[j].ID
	})

	return pokedex

}
