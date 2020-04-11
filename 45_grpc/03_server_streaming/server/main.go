package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/dimzrio/grpc-courses/03_server_streaming/model"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

type server struct {
}

func (*server) DetailsPerson(req *model.GreetingReq, stream model.NameService_DetailsPersonServer) error {

	log.Println("Server streaming starting with request", req)

	name := req.GetName()

	for i := 1; i <= 10; i++ {
		result := fmt.Sprintf("Resp-%d => Hello %s", i, name)

		resp := &model.GreetingResp{
			Result: result,
		}

		err := stream.Send(resp)
		time.Sleep(1 * time.Second)

		if err != nil {
			log.Fatal(err)
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
	model.RegisterNameServiceServer(gRPCServer, &server{})

	err = gRPCServer.Serve(listen)

	if err != nil {
		log.Fatal(err)
	}

}
