package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"strconv"
	"time"
)

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
	var args Arguments
	var reply Reply

	ch := make(chan string)

	// compatible to receive loop count from os env.
	count := flag.String("count", "100", "loop max amount")
	flag.Parse()

	loopCount, _ := strconv.Atoi(*count)

	for i := 0; i < loopCount; i++ {
		go func() {
			randomCaller(args, reply, ch)
		}()
		fmt.Println(<-ch)
		fmt.Println(i)
	}

	close(ch)
}

func randomCaller(args Arguments, reply Reply, ch chan string) {
	rand.Seed(time.Now().UnixNano())
	args.FirstParam = rand.Intn(10)
	args.SecondParam = rand.Intn(10)

	if caller("RpcServer.Multiply", args, &reply) {
		ch <- fmt.Sprintf(
			"%d result %d and %d is: %d\n",
			time.Now().UnixNano(),
			args.FirstParam,
			args.SecondParam,
			reply.Result,
		)
	}

}

func caller(rpcMethod string, args interface{}, reply interface{}) (res bool) {
	fmt.Println("Making Request")

	dial, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("Dial err:", err)
		return false
	}

	defer dial.Close()

	err = dial.Call(rpcMethod, args, reply)
	if err != nil {
		log.Fatal("RPC call err:", err)
		return false
	}

	return true
}
