package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/dimzrio/grpc-courses/example/prime/model"

	"google.golang.org/grpc"
)

const serverAddr = "localhost:50051"

func main() {
	fmt.Println("*** Client ***")

	// open conn
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
		return
	}

	// close conn
	defer conn.Close()

	// open grpc
	gRPC := model.NewPrimeServiceClient(conn)

	// open req
	req := &model.PrimeReq{
		Number: 100,
	}

	resp, err := gRPC.PrimeDecompotion(context.Background(), req)

	if err != nil {
		log.Fatalf("Found error at send req: %v", err)
	}

	for {
		msg, err := resp.Recv()

		if err == io.EOF {
			log.Println("Close by end of stream response")
			break
		}

		if err != nil {
			log.Printf("Found error at reading resp: %v", err)
			break
		}

		log.Println(msg.GetResult())
	}
}
