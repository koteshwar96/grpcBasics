package main

import (
	"fmt"
	"grpcLearning/grpc/server"
)

func main() {
	fmt.Println("Hello world")
	server.StartUnaryServer()
}
