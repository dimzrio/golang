package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/dimzrio/grpc-courses/example/maxNum/model"

	"google.golang.org/grpc"
)

func main() {
	log.Println("*** Bi-Directional Client ***")
	var payload []*model.MaxNumReq

	var data = []int32{1, 5, 3, 6, 2, 20}

	// Generate payload
	for _, each := range data {
		req := &model.MaxNumReq{
			Number: each,
		}

		payload = append(payload, req)
	}

	// open conn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Can't connect to server: %v", err)
	}

	defer conn.Close()

	gRPC := model.NewMaxNumberClient(conn)

	// Create stream invoking by stream
	stream, err := gRPC.GetMax(context.Background())

	if err != nil {
		log.Fatalf("Failed to call service GetExponent: %v", err)
	}

	done := make(chan bool)

	// Send message
	go func() {
		for _, data := range payload {
			log.Printf("Send data -> %d", data.GetNumber())

			err = stream.Send(data)
			if err != nil {
				log.Fatalf("Failed send data: %v", data)
			}
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	// Recive message
	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Found error at reading resp: %v", err)
				break
			}

			log.Printf("Max Number: %d", resp.GetResult())
		}
		close(done)
	}()

	// Block until everything done
	<-done
}
