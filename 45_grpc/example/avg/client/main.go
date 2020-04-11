package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dimzrio/grpc-courses/example/avg/model"

	"google.golang.org/grpc"
)

const serverAddr = "localhost:50051"

func main() {
	log.Println("*** Client Stream ***")

	var sendData []*model.AvgReq

	// Generate requst
	avgData := []int64{1, 2, 3, 4}
	for _, each := range avgData {

		req := &model.AvgReq{
			Number: each,
		}

		sendData = append(sendData, req)
	}

	// Open conn
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("Found error at connect to server: %v", err)
	}

	gRPC := model.NewAvgServiceClient(conn)

	// Call service sumAvg
	stream, err := gRPC.SumAvg(context.Background())

	if err != nil {
		log.Fatalf("Found error at call service sumAvg: %v\n", err)
	}

	// Send data
	for _, each := range sendData {
		log.Println("Send data => ", each.GetNumber())

		err = stream.Send(each)

		if err != nil {
			log.Fatalf("Found error at send requst: %v", err)
		}

		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Found error at read response: %v", err)
	}

	fmt.Println("Avg:", resp.GetResult())

}
