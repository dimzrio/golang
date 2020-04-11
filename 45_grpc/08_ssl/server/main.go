package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc/status"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"

	"github.com/dimzrio/grpc-courses/07_deadlines/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) SayHello(ctx context.Context, req *model.HelloReq) (*model.HelloResp, error) {

	// If client enable timeout
	if ctx.Err() == context.Canceled {
		log.Println("Request was cancelled by timeout")
		return nil, status.Error(codes.Canceled, "The client cancelled")
	}

	firstname := req.GetFirstname()
	lastname := req.GetLastname()

	msg := fmt.Sprintf("Hello, %s %s", firstname, lastname)

	resp := &model.HelloResp{
		Result: msg,
	}

	log.Println(req)
	return resp, nil
}

func main() {
	log.Println("*** Unary with deadlines ***")

	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Printf("Found error: %v\n", err)
	}

	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pam"
	cred, err := credentials.NewServerTLSFromFile(certFile, keyFile)

	if err != nil {
		log.Printf("Found error ssl: %v\n", err)
	}

	gRPCServer := grpc.NewServer(grpc.Creds(cred))
	model.RegisterHelloServiceServer(gRPCServer, &server{})

	gRPCServer.Serve(listen)
}
