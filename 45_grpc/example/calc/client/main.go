package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dimzrio/grpc-courses/calc/model"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("*** Unary Calc Client ***")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Printf("Found error: %v", err)
	}

	defer conn.Close()

	// Open gRPC
	gRPC := model.NewCalcSumClient(conn)

	// Set Payload
	req := &model.RequestSum{
		X: 7,
		Y: 3,
	}

	// Call Service
	resp, err := gRPC.Sum(context.Background(), req)

	if err != nil {
		log.Printf("Found error: %v", err)
	}

	fmt.Println("Response from server:", resp.Result)

}
