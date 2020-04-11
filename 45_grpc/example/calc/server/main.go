package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dimzrio/grpc-courses/calc/model"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *model.RequestSum) (*model.ResponseSum, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y

	resp := &model.ResponseSum{
		Result: result,
	}

	log.Printf("Request from client: %v", req)
	return resp, nil
}

func main() {

	fmt.Println("*** Unary Calc Server ***")

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Found error: %v\n", err)
	}

	serv := grpc.NewServer()
	model.RegisterCalcSumServer(serv, &server{})

	err = serv.Serve(lis)
	if err != nil {
		log.Fatalf("Found error: %v\n", err)
	}

}
