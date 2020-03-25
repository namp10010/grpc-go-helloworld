package main

import (
    "context"
    "log"
    "net"

    pb "../proto"
    grpc "google.golang.org/grpc"
)

const (
    port = ":50051"
)

type greeterServer struct {
}

func (s *greeterServer) Greet(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    log.Printf("Received request: %v", in.GetName())
       return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
    tcpListener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()

       pb.RegisterGreeterServer(grpcServer, &greeterServer{})
       if err := grpcServer.Serve(tcpListener); err != nil {
           log.Fatalf("failed to serve: %v", err)
       }
}
