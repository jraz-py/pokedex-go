package pokemon

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var mockServer *httptest.Server

func init() {
	response := `{
		"id": 1,
		"name": "bulbasaur",
		"weight": 200,
		"types": [
			{
				"type": {
					"name": "grass"
				}
			}
		],
		"moves": [
			{
				"move": {
					"name": "razor-wind"
				}
			}
		]
	}`
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}))
}

func Test_getPokemon(t *testing.T) {

	tests := []struct {
		name string
		want Pokemon
	}{
		{
			name: "testOne",
			want: Pokemon{
				ID:     1,
				Name:   "bulbasaur",
				Wieght: 200,
				Type: []Types{
					{
						ElementType: ElemType{
							Name: "grass",
						},
					},
				},
				AttackMoves: []Moves{
					{
						Move: Move{
							Name: "razor-wind",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := getPokemon(mockServer.URL); !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("getPokemon(%v) => got %v, want %v ", mockServer.URL, got, tt.want)
			}
		})
	}
}

func Test_BuildPokedex(t *testing.T) {

	pokeToGet := 2
	sem := NewSem(2)
	tests := []struct {
		name string
		want []Pokemon
	}{
		{
			name: "testOne",
			want: []Pokemon{
				{
					ID:     1,
					Name:   "bulbasaur",
					Wieght: 200,
					Type: []Types{
						{
							ElementType: ElemType{
								Name: "grass",
							},
						},
					},
					AttackMoves: []Moves{
						{
							Move: Move{
								Name: "razor-wind",
							},
						},
					},
				},
				{
					ID:     1,
					Name:   "bulbasaur",
					Wieght: 200,
					Type: []Types{
						{
							ElementType: ElemType{
								Name: "grass",
							},
						},
					},
					AttackMoves: []Moves{
						{
							Move: Move{
								Name: "razor-wind",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := BuildPokedex(sem, pokeToGet, mockServer.URL); !equal(got, tt.want) {
				t.Fatalf("BuildPokedex() => got %v, want %v ", got, tt.want)
			}
		})
	}
}

func equal(a, b []Pokemon) bool {
	fmt.Printf("a %v\nb %v", a, b)

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.String() != b[i].String() {
			return false
		}
	}
	return true
}
