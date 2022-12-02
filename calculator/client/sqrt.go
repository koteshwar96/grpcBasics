package main

import (
	"context"
	"fmt"
	pb "grpcLearning/calculator/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func calculateSqRt(client pb.CalculatorServiceClient, number int ){

	log.Printf("Making rpc call to calculate sqrt of %d ",number)

	response,err := client.CalculateSqRt(context.Background(), &pb.SqRtRequest{Number: int32(number)})

	if err != nil {
		//checking if this is a known/predefined error
		e, ok := status.FromError(err)

		// if this is a grpc error then it would be true
		if ok {
			fmt.Errorf("server has responded with an error. Code is: %d, message is %s",e.Code(),e.Message())

			// Depending on the error error code one can act accordingly or build business logic around it
			if e.Code() == codes.InvalidArgument {
				log.Print("Proably we might have sent an invalid data for rpc call")
			}
		} else {
			fmt.Errorf("Error occured while making rpc call to determine sqrt of a number. Error: %s",err)
		}
		return
	}

	log.Printf("Received sqrt of %d and the result is %f",number,response.GetResult())

}