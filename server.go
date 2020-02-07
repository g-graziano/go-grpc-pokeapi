package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/g-graziano/go-grpc-pokeapi/proto"
	pokeapi "github.com/g-graziano/go-grpc-pokeapi/resource/pokemon"
	"google.golang.org/grpc"
)

const port string = "50051"

type server struct {
	pb.UnimplementedPokemonServer
	pokemon *pb.GetPokemonResponse
}

func (s *server) GetPokemon(stream pb.Pokemon_GetPokemonServer) error {

}

func main() {
	lis, err := net.Listen("tcp", "localhost:"+port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPokemonServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
