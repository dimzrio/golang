package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dimzrio/grpc-courses/02_unary/model"

	"google.golang.org/grpc"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) SayHello(ctx context.Context, req *model.RequestName) (*model.ResponseName, error) {
	firstName := req.GetName().GetFirstname()
	lastName := req.GetName().GetLastname()

	result := fmt.Sprintf("Hello %s %s..!!!", firstName, lastName)

	resp := &model.ResponseName{
		Result: result,
	}
	fmt.Println(req)
	return resp, nil
}

func main() {
	fmt.Println("gRPC Unary")

	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	model.RegisterNameServiceServer(gRPCServer, &server{})

	err = gRPCServer.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}

}
