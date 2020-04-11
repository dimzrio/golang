package main

import (
	"context"
	"fmt"
	"log"

	greetpb "github.com/dimzrio/grpc-courses/greet/greatpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("*** Client Unary ***")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	c := greetpb.NewGreatServiceClient(cc)

	// Unary
	doUnary(c)

}

func doUnary(c greetpb.GreatServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Dimas",
			LastName:  "Riotantowi",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response from server: %v", res.Result)
}
