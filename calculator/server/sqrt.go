package main

import (
	"context"
	"fmt"
	pb "grpcLearning/calculator/proto"
	"log"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CalculateSqRt(context context.Context, request *pb.SqRtRequest) (*pb.SqRtResponse, error) {

	log.Printf("Received request to calculate sqrt of %d ",request.GetNumber())

	if(request.GetNumber() <0 ){
		log.Print("Failing the request as we have received a negative number to calculate sqrt")

		// we are return a custom grpc error so that client can appropriate actions like retry if needed
		// We can add more error details, refer here https://cloud.google.com/apis/design/errors
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received negative number %d",request.GetNumber())) 
	}

	return &pb.SqRtResponse{
		Result: math.Sqrt(float64(request.GetNumber())),
	},nil
	
}