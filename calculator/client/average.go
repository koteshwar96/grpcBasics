package main

import (
	"context"
	pb "grpcLearning/calculator/proto"
	"log"
	"time"
)
func calculateAverage (grpcClient  pb.CalculatorServiceClient) {
	log.Print("making gRPC average request ")

	readings := []*pb.AverageRequest {
		{ReadingValue: 6},
		{ReadingValue: 8},
		{ReadingValue: 46},
		{ReadingValue: 61},
	}

	averageClient,err := grpcClient.CalculateAverage(context.Background())

	if err!= nil {
		log.Fatal("Error while making average gRPC call. Error: ",err)
	}

	for _,reading := range readings {
		averageClient.Send(reading)
		log.Print("waiting for 1 sec before we send next data to server...")
		time.Sleep(1*time.Second)
	}
	log.Print("Completed sending data to the server...")
	
	averageResponse, err := averageClient.CloseAndRecv()

	if err != nil {
		log.Fatal("Error while receiving average gRPC call response. Error: ",err)
	}

	log.Printf("Average of the readings is: %f ",averageResponse.GetReadingAverage())

}