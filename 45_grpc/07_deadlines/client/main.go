package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/dimzrio/grpc-courses/07_deadlines/model"

	"google.golang.org/grpc"
)

func main() {
	log.Println("*** Client unary with deadlines ***")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Found Error : %v\n", err)
	}

	defer conn.Close()

	req := &model.HelloReq{
		Firstname: "Dimas",
		Lastname:  "Rio",
	}

	gRPC := model.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // Timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Success
	defer cancel()

	resp, err := gRPC.SayHello(ctx, req)

	if err != nil {
		// ok default is false
		statusErr, ok := status.FromError(err)

		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Timeout Hit!")
			} else {
				log.Printf("Found Unexpected Error: %v", statusErr.Message())
			}
		} else {
			log.Println("Error at callinng SayHello GRPC")
		}
	}

	log.Println(resp)

}
