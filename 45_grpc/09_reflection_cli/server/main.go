package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/dimzrio/grpc-courses/09_reflection_cli/model"
)

type server struct {
}

var result float64

func (*server) AdditionService(ctx context.Context, req *model.CalcReq) (*model.CalcResp, error) {
	number1 := req.GetFirstNumber()
	number2 := req.GetSeconNumber()

	resp := &model.CalcResp{
		Result: number1 + number2,
	}

	return resp, nil
}

func (*server) SubtractionService(ctx context.Context, req *model.CalcReq) (*model.CalcResp, error) {
	number1 := req.GetFirstNumber()
	number2 := req.GetSeconNumber()

	resp := &model.CalcResp{
		Result: number1 - number2,
	}

	return resp, nil

}

func (*server) MultiplicationService(ctx context.Context, req *model.CalcReq) (*model.CalcResp, error) {

	number1 := req.GetFirstNumber()
	number2 := req.GetSeconNumber()

	resp := &model.CalcResp{
		Result: number1 * number2,
	}

	return resp, nil

}

func (*server) DivisionService(ctx context.Context, req *model.CalcReq) (*model.CalcResp, error) {

	number1 := req.GetFirstNumber()
	number2 := req.GetSeconNumber()

	resp := &model.CalcResp{
		Result: number1 / number2,
	}

	return resp, nil

}

func main() {

	log.Println("*** Calculator Service ***")

	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Found error at open port: %v\n", err)
	}

	gRPC := grpc.NewServer()
	model.RegisterCalcatorServiceServer(gRPC, &server{})
	reflection.Register(gRPC)
	gRPC.Serve(listen)
}

// Connect with evans cli
// brew install evans
// https://github.com/ktr0731/evans

// $ evans -r

// default@127.0.0.1:50051> desc calcReq
// +-------------+-------------+
// |    FIELD    |    TYPE     |
// +-------------+-------------+
// | firstNumber | TYPE_DOUBLE |
// | seconNumber | TYPE_DOUBLE |
// +-------------+-------------+

// default@127.0.0.1:50051> show service
// +-----------------+-----------------------+-------------+--------------+
// |     SERVICE     |          RPC          | REQUESTTYPE | RESPONSETYPE |
// +-----------------+-----------------------+-------------+--------------+
// | CalcatorService | AdditionService       | calcReq     | calcResp     |
// |                 | SubtractionService    | calcReq     | calcResp     |
// |                 | MultiplicationService | calcReq     | calcResp     |
// |                 | DivisionService       | calcReq     | calcResp     |
// +-----------------+-----------------------+-------------+--------------+

// default@127.0.0.1:50051> service CalcatorService

// default.CalcatorService@127.0.0.1:50051> call AdditionService
// firstNumber (TYPE_DOUBLE) => 5
// seconNumber (TYPE_DOUBLE) => 10
// {
//   "result": 15
// }

// default.CalcatorService@127.0.0.1:50051> call DivisionService
// firstNumber (TYPE_DOUBLE) => 10
// seconNumber (TYPE_DOUBLE) => 3
// {
//   "result": 3.3333333333333335
// }

// default.CalcatorService@127.0.0.1:50051>
