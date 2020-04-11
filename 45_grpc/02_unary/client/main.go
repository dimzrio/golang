package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dimzrio/grpc-courses/02_unary/model"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("*** Client ***")

	// Open Connection
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	// Close Connection
	defer connection.Close()

	// open grpc
	gRPC := model.NewNameServiceClient(connection)

	// Create req
	req := &model.RequestName{
		Name: &model.Name{
			Firstname: "Dimas",
			Lastname:  "Riotantowi",
		},
	}

	resp, err := gRPC.SayHello(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.GetResult())

}
