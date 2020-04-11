package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	greetpb "github.com/dimzrio/grpc-courses/greet/greatpb"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	result := fmt.Sprintf("Hello %s %s\n", firstname, lastname)
	resp := &greetpb.GreetResponse{
		Result: result,
	}
	log.Println(req)
	return resp, nil
}

func main() {
	fmt.Println("*** Server ***")

	lis, err := net.Listen("tcp", ":50051")

	if err != err {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreatServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
