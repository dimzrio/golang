package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/dimzrio/grpc-courses/05_bi_directional_streaming/model"

	"google.golang.org/grpc"
)

func main() {
	log.Println("*** Bi-Directional Client ***")
	var payload []*model.ExponentReq

	// Generate payload
	for i := 1; i <= 10; i++ {
		req := &model.ExponentReq{
			Number: int32(i),
		}

		payload = append(payload, req)
	}

	// open conn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Can't connect to server: %v", err)
	}

	defer conn.Close()

	gRPC := model.NewExponent2Client(conn)

	// Create stream invoking by stream
	stream, err := gRPC.GetExponent(context.Background())

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

			log.Printf("Respond: Exponent2 from number %d is %d", resp.GetNumber(), resp.GetResult())
		}
		close(done)
	}()

	// Block until everything done
	<-done
}
