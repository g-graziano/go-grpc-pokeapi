package main

import (
	"context"
	"log"

	pb "github.com/g-graziano/go-grpc-pokeapi/proto"
	"google.golang.org/grpc"
)

func getPokemonFromServer(stream pb.Pokemon_GetPokemonClient) {

}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create pokemon stream
	client := pb.NewPokemonClient(conn)
	stream, err := client.GetPokemon(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	getPokemonFromServer(stream)
}
