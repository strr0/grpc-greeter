package main

import (
	"context"
	"google.golang.org/grpc"
	"greeter/proto"
	"log"
	"net"
)

type Say struct {
	greeter.UnimplementedSayServer
}

func (Say) Hello(ctx context.Context, req *greeter.Request) (*greeter.Response, error) {
	log.Println("Received Say.Hello request")
	return &greeter.Response{
		Msg: "Hello " + req.Name,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	srv := grpc.NewServer()
	greeter.RegisterSayServer(srv, new(Say))
	if err := srv.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
