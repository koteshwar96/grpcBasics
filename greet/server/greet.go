package main

import (
	"context"
	"fmt"
	pb "grpcLearning/greet/proto"
)


func (s *Server) Greet(ctx context.Context,request *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Request received to greet %s",request.GetName())

	return &pb.GreetResponse{
		Response: fmt.Sprintf("Hi %s, How are you doing ?",request.GetName()),
	},nil
}
