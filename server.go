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
	log.Println("start new server")
	ctx := stream.Context()

	for {
		// exit if context is done
		select {
		case <-ctx.Done():
			log.Printf("signal interrupt")
			return ctx.Err()
		default:
		}

		req, err := stream.Recv()
		if err == io.EOF {
			// return will close stream from server
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		fmt.Println("receive ", req)

		s.pokemon, err = pokeapi.GetPokemon(req)

		fmt.Println("Pokemon send: ", s.pokemon)

		stream.Send(s.pokemon)
	}
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
