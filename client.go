package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/g-graziano/go-grpc-pokeapi/proto"
	"google.golang.org/grpc"
)

func getPokemonFromServer(stream pb.Pokemon_GetPokemonClient) {
	rand.Seed(time.Now().UnixNano())
	ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random numbers to pokemon stream
	go func() {
		for i := 1; i <= 10; i++ {
			randomID := int32(rand.Intn(300))

			req := pb.GetPokemonInput{Id: randomID}

			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}

			fmt.Println("send pokemon id: " + fmt.Sprint(randomID))

			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				close(done)
				return
			}

			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			fmt.Println("Pokemon received:", resp)
		}
	}()

	// third goroutine closes done channel
	go func() {
		<-ctx.Done()
		log.Printf("exit stream")
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Printf("finished stream")
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
