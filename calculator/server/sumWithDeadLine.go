package main

import (
	"context"
	"fmt"
	pb "grpcLearning/calculator/proto"
	"time"
)
func (s *Server) CalculateSumWithDeadLine(ctx context.Context, request *pb.SumRequest) (*pb.SumResponse, error) {

	// Here we are using context to get the value/data
	fmt.Printf("Received CalculateSumWithDeadLine request to server, Request is: %s, traceId: %s \n", request, ctx.Value("myKey") )

	// yet to figure out why we are unable to read values from the context
	fmt.Println("traceid is  from ctx: ",ctx.Value("myKey"))

	for i:=0; i<5; i++ {
        time.Sleep(1*time.Second)
        fmt.Println(i+1," Sec completed...")
    }
	fmt.Println("waiting completed and sending response now...")
	
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Need not calulate response value as we have crossed deadline already")
		return nil,ctx.Err()
	}

	return &pb.SumResponse{
		Sum: request.FirstNumber + request.SecondNumber,
	},nil
}