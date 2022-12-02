package main

import (
	"context"
	pb "grpcLearning/calculator/proto"
	"log"
)

func calculateSum(grpcClient pb.CalculatorServiceClient){
	log.Print("Now making gRPC call to calculate sum")

	firstNumber := 3
	secondNumber := 66
	request := &pb.SumRequest{
		FirstNumber: int32(firstNumber),
		SecondNumber: int32(secondNumber),
	}
	
	response, err := grpcClient.CalculateSum(context.Background(), request)
	if err != nil {
		log.Println("Error while making RPC call. Error:",err)
	}
	
	log.Printf("Sum of %d and %d is %d",firstNumber,secondNumber, response.GetSum())
}