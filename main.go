package main

import (
	"log"
	"net"

	"github.com/NothingXiang/go-rpc/api"
	"github.com/NothingXiang/go-rpc/server"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":2345")

	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(nil)

	api.RegisterMyServiceServer(grpcServer, &server.Server{})

	grpcServer.Serve(listener)

}
