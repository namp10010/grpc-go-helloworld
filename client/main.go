package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "../proto"
	grpc "google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "nam"
)

func main() {
	grpcConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer grpcConn.Close()
	greeterClient := pb.NewGreeterClient(grpcConn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := greeterClient.Greet(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not great: %v", err)
	}
	log.Printf("Received response: %s", response.GetMessage())
}
