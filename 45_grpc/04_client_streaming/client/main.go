package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dimzrio/grpc-courses/04_client_streaming/model"

	"google.golang.org/grpc"
)

const serverAddr = "localhost:50051"

func main() {
	fmt.Println("*** Client Stream ***")

	var sumData []*model.SumReq

	for i := 1; i <= 10; i++ {
		req := &model.SumReq{
			Number: int64(i),
		}

		sumData = append(sumData, req)
	}

	// open conn
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Found error connect to server: %v\n", err)
	}

	defer conn.Close()

	gRPC := model.NewSumNumberClient(conn)

	stream, err := gRPC.SumAll(context.Background())

	if err != nil {
		log.Fatalf("Found error at call service sumAll: %v\n", err)
	}

	// Send stream
	for _, each := range sumData {
		log.Printf("Sending request -> %v \n", each)
		err = stream.Send(each)

		if err != nil {
			log.Fatalf("Found error at sending req: %v", err)
		}

		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Found error at read response: %v", err)
	}

	fmt.Println("Total Sum:", resp.GetResult())
}
