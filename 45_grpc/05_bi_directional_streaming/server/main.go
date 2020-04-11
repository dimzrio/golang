package main

import (
	"io"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	"github.com/dimzrio/grpc-courses/05_bi_directional_streaming/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) GetExponent(stream model.Exponent2_GetExponentServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Found error at reading request: %v", req)
		}

		number := req.GetNumber()
		result := &model.ExponentResp{
			Result: int32(math.Exp2(float64(number))),
			Number: int32(number),
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
	model.RegisterExponent2Server(gRPCServer, &server{})

	// Server grpc
	gRPCServer.Serve(listen)
}
