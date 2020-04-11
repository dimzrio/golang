package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/dimzrio/grpc-courses/example/prime/model"

	"google.golang.org/grpc"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) PrimeDecompotion(req *model.PrimeReq, stream model.PrimeService_PrimeDecompotionServer) error {

	log.Println("Server streaming starting with request", req)

	number := req.GetNumber()

	var divisor int64 = 2

	for number > 1 {
		if number%divisor == 0 {
			resp := &model.PrimeResp{
				Result: divisor,
			}

			err := stream.Send(resp)
			time.Sleep(1 * time.Second)

			if err != nil {
				log.Fatal(err)
			}

			number = number / divisor
		} else {
			divisor++
			log.Printf("Divisor has increader to -> %d\n", divisor)
		}
	}
	return nil

}
func main() {

	fmt.Println("gRCP Server Streaming")

	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	model.RegisterPrimeServiceServer(gRPCServer, &server{})

	err = gRPCServer.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}

}
