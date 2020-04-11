package main

import (
	"context"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dimzrio/grpc-courses/06_error_handling/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) SquareRoot(ctx context.Context, req *model.SquareReq) (*model.SquareResp, error) {
	number := float64(req.GetNumber())

	log.Printf("Requsts from client: %v\n", req)

	if number < 1 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Can't using negative number")
	}

	resp := &model.SquareResp{
		Result: math.Sqrt(number),
	}

	return resp, nil
}

func main() {
	log.Println("*** Unary Error Handling Example ***")

	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatal(err)
	}

	gRPCServer := grpc.NewServer()
	model.RegisterSquareServiceServer(gRPCServer, &server{})

	err = gRPCServer.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}
}
