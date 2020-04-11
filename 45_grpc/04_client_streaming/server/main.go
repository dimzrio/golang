package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dimzrio/grpc-courses/04_client_streaming/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) SumAll(stream model.SumNumber_SumAllServer) error {

	var sum int64

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("Close stream by client")

			resp := &model.SumResp{
				Result: sum,
			}

			return stream.SendAndClose(resp)
		}

		if err != nil {
			log.Printf("Found error at receive request stream: %v", err)
		}

		number := req.GetNumber()

		sum = sum + number
	}
}

func main() {
	fmt.Println("*** gRPC Client Stream ***")
	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatalf("Found error at listen port: %v\n", err)
	}

	gRCPServer := grpc.NewServer()
	model.RegisterSumNumberServer(gRCPServer, &server{})

	err = gRCPServer.Serve(listen)

	if err != nil {
		log.Fatalf("Found error at serve grpc: %v", err)
	}

}
