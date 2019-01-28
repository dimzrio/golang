package main

import (
	"log"
	"net"

	"github.com/novalagung/gubrak"
	"google.golang.org/grpc"
)

const (
	protocol = "tcp"
	port     = ":50051"
)

func main() {

	listen, err := net.Listen(protocol, port)

	if err != nil {
		log.Fatalf("Found error: %v", err)
		return
	}

	log.Println("Starting server...")
	log.Println("Random digit:", gubrak.RandomInt(10, 20))
	grpc := grpc.NewServer()
	grpc.Serve(listen)
}
