package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dimzrio/grpc-courses/example/maxNum/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) GetMax(stream model.MaxNumber_GetMaxServer) error {
	var temp int32
	temp = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Found error at reading request: %v", req)
		}

		number := req.GetNumber()
		result := &model.MaxNumResp{
			Result: func() int32 {
				if temp < number {
					temp = number
				}
				return temp
			}(),
		}

		log.Println(req)

		err = stream.Send(result)
		if err != nil {
			log.Fatalf("Found error at sending response: %v", err)
		}
	}
}

func main() {
	log.Println("*** Bi-Directional Stream ***")

	listen, err := net.Listen(protocol, port)
	defer listen.Close()

	if err != nil {
		log.Fatalf("Can't open port: %v", err)
	}

	// Create server grpc
	gRPCServer := grpc.NewServer()

	// Register to protobuf
	model.RegisterMaxNumberServer(gRPCServer, &server{})

	// Server grpc
	gRPCServer.Serve(listen)
}
