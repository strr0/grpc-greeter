package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"greeter/proto"
	"log"
)

func main() {
	dial, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	cli := greeter.NewSayClient(dial)
	res, err := cli.Hello(context.Background(), &greeter.Request{Name: "John"})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
