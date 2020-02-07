package pokeapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/g-graziano/go-grpc-pokeapi/proto"
)

type Pokeapi interface {
	GetPokemon(input *pb.GetPokemonInput) (*pb.GetPokemonResponse, error)
}

func GetPokemon(input *pb.GetPokemonInput) (*pb.GetPokemonResponse, error) {
	var pokemon pb.GetPokemonResponse

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + fmt.Sprint(input.Id))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.Unmarshal(contents, &pokemon)

	return &pokemon, nil
}
