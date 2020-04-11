package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/dimzrio/grpc-courses/06_error_handling/model"
	"google.golang.org/grpc"
)

func doUnary(u model.SquareServiceClient, number int32) {
	resp, err := u.SquareRoot(context.Background(), &model.SquareReq{
		Number: number,
	})

	if err != nil {
		respErr, ok := status.FromError(err)

		// ok -> if error status
		if ok {
			fmt.Println(respErr.Code())
			fmt.Println(respErr.Message())

			if respErr.Code() == codes.InvalidArgument {
				log.Println("Please make sure, don't input negative number")
				return
			}

		} else {
			log.Fatalf("Error at calling service: %v\n", err)
			return
		}
	}

	log.Printf("Response: %v\n", float64(resp.GetResult()))
}

func main() {
	fmt.Println("*** Client ***")

	// Open Connection
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer connection.Close()

	if err != nil {
		log.Fatal(err)
	}

	// open grpc
	gRPC := model.NewSquareServiceClient(connection)

	// Call doUnary correct result
	doUnary(gRPC, 3)

	// Call doUnary incorrect result
	doUnary(gRPC, -1)
}
