package main

import (
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dimzrio/grpc-courses/example/avg/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) SumAvg(stream model.AvgService_SumAvgServer) error {

	var temp int64
	var div int64
	var avg float32

	for {

		// Read client request stream
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("Closed by client stream EOF")

			avg = float32(temp) / float32(div)

			resp := &model.AvgResp{
				Result: avg,
			}

			return stream.SendAndClose(resp)

		}

		if err != nil {
			log.Fatalf("Found error at read requsts: %v", err)
		}

		number := req.GetNumber()
		log.Println("Recv number ->", number)

		temp = temp + number
		div++
	}
}

func main() {
	log.Println("*** gRPC Client Stream ***")

	// open conn
	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatalf("Found error at open connection: %v", err)
	}

	gRPCServer := grpc.NewServer()
	model.RegisterAvgServiceServer(gRPCServer, &server{})

	err = gRPCServer.Serve(listen)

	if err != nil {
		log.Fatalf("Found error at listen port: %v", err)
	}

}
