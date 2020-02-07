package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	pb "https://github.com/g-graziano/go-grpc-pokeapi/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
}
