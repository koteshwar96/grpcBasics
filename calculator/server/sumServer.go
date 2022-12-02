package main

import (
	"context"
	pb "grpcLearning/calculator/proto"
	"log"
)


func (server *Server) CalculateSum(ctx context.Context,request *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Came to calculate sum method with request : %S",request.String())
	
	return &pb.SumResponse{
		Sum: request.GetFirstNumber() + request.GetSecondNumber(),
	},nil
}