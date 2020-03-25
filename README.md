# Go gRPC Hello World

This is a clone from official [gRPC Go's Quickstart](https://grpc.io/docs/quickstart/go/) with minor changes. Majority of variables are also renamed so that they make more sense for beginners.

## Build

### Prerequisites

* go 1.14 or later (shouldn't matter much but just in case)
* libprotoc 3.11.4 (for protoc)

### Generate protobuf service definition

```zsh
protoc -I proto proto/hello.proto --go_out=plugins=grpc:proto
```

A new go file should be generated as ```proto/hello.pb.go``` that has the service definition in Go.

### Create server

The folowing code in ```server/main.go``` implements the server interface to serve the request.

```go
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

```

Please note that this server ```type greeterServer struct {}``` doesn't have an ```pb.UnimplementedGreeterServer``` as in the reference document. This is because I don't understand why it's required when ```greetServer``` implements type ```hello.GreetServer``` anyway.

Compile the server code

```zsh
go build -o server/main server/main.go
```

### Create client

The folowing code in ```client/main.go``` implements the server interface to serve the request.

```go
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
```

The client takes one command line argument as a name to pass to the request to server or use a default if one is missing. Compile the code.

```zsh
go build -o client/main client/main.go
```

## Run

Run the server (here from zsh)

```zsh
./server/main
```

On another terminal run the client passing in an argument

```zsh
./client/main world
```

On the server side there should be

```zsh
2020/03/25 22:18:20 Received request: world
```

And on client side

```zsh
2020/03/25 22:18:20 Received response: Hello world
```

To exit the server press Ctrl+C.

## Reference

* [gRPC Go Quickstart](https://grpc.io/docs/quickstart/go/)
