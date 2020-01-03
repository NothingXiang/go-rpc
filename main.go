package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/NothingXiang/go-rpc/server"
)

func main() {

	// 注册http rpc服务
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	// 监听socket
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(listen, nil)

	// 等待运行
	time.Sleep(2 * time.Second)

	// 客户端调用
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}

// colon