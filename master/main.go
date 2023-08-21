package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type RpcServer struct{}

type (
	Arguments struct {
		FirstParam  int
		SecondParam int
	}

	Reply struct {
		Result int
	}
)

func main() {
	master := new(RpcServer)
	master.server()
}

func (a *RpcServer) server() {
	if err := rpc.Register(a); err != nil {
		log.Panic(err)
	}

	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatal(fmt.Sprintf("Listen err: %s", err.Error()))
	}

	defer listener.Close()

	fmt.Println("RPC Listening...")

	for {
		rpcConn, listenerErr := listener.Accept()

		if listenerErr != nil {
			continue
		}

		go rpc.ServeConn(rpcConn)
	}
}

func (a *RpcServer) Multiply(args Arguments, res *Reply) (err error) {
	fmt.Printf("%d numbers: %d and %d\n", time.Now().UnixNano(), args.FirstParam, args.SecondParam)
	res.Result = args.FirstParam * args.SecondParam
	return
}
