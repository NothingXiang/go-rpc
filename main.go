package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/NothingXiang/go-rpc/api"
	"github.com/NothingXiang/go-rpc/server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":2345"
	gatePort = ":2346"
)

func main() {

	// 注册grpc-gateway服务
	go RegisterGateway(fmt.Sprintf("localhost%v", grpcPort), gatePort)

	// 注册grpc 服务（阻塞的）
	RegisterGrpc(grpcPort)

}

func RegisterGrpc(port string) {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterMyServiceServer(grpcServer, &server.Server{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

// 注册一个gateway服务
func RegisterGateway(grpcAddr, gatewayPort string) {

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	api.RegisterMyServiceHandlerFromEndpoint(context.TODO(), mux, grpcAddr, opts)

	if err := http.ListenAndServe(gatewayPort, mux); err != nil {
		log.Fatalln(err)
	}
}
