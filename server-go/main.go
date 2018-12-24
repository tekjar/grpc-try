package main

//protoc helloworld.proto --go_out=plugins=grpc:.
import (
	"context"
	"log"
	"net"
	"server/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.Data, error) {
	data := helloworld.Data{
		A: "ather",
		B: "pure",
		C: "grpc",
		D: "load",
		E: "testing",
		F: "with",
		G: "go",
		H: "lang",
		I: 1.0,
		J: 2.0,
		K: 3.0,
		L: 4.0,
		M: 5.0,
		N: 6.0,
		O: 7.0,
		P: 8.0,
		Q: 1,
		R: 2,
		S: 3,
		T: 4,
		U: 5,
		V: 6,
		W: 7,
		X: 8,
	}

	return &data, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
