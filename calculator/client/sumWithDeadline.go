package main

import (
	"context"
	"fmt"
	pb "grpcLearning/calculator/proto"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)
func calculateSumWithDeadline(client pb.CalculatorServiceClient, timeout time.Duration){
	log.Printf("Making calculateSumWithDeadline grpc call with timeout of %d ",timeout)

	var firstNumber int32 = 33
	var secondNumber int32 = 77

	// we are creating basic or empty context
	ctx := context.Background();

	var traceId = uuid.New().String()
	// we are adding data for the context that can be used by server 
	ctx = context.WithValue(ctx,"myKey",traceId)
	fmt.Println("traceid is  from ctx: ",ctx.Value("myKey"))

	// adding timeout to the context, cancelFunc will be automatically called post timeout, if required we can call before that as well using cancelFunction
	ctx, cancelFunction := context.WithTimeout(ctx,timeout)

	// we are calling cancel function before we exit from this function just to make sure we cleanup
	defer cancelFunction()

	//As soon as the timeour occurs we no longer wait here for the response and we will get grpc err with context deadline exceeded
	sumResponse, err :=  client.CalculateSumWithDeadLine(ctx,&pb.SumRequest{
		FirstNumber: firstNumber,
		SecondNumber: secondNumber,
	})

	if err != nil {
		// first check if this is a grpc error
		e,ok := status.FromError(err)
		if (ok) {
			// we can add our business logic like retry depending on the error code or message
			fmt.Printf("calculateSumWithDeadline failed with grpc error. Error code: %d, Error description: %s ",e.Code(),e.Message())
		} else {
			fmt.Printf("Error while making calculateSumWithDeadline grpc call. Error: %s",err)
		}
		
		return
	}

	log.Printf("calculateSumWithDeadline response is: %d",sumResponse.GetSum())

}