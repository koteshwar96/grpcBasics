package main

import (
	"context"
	pb "grpcLearning/calculator/proto"
	"io"
	"log"
	"time"
)

func makeMaxRPCCall(client pb.CalculatorServiceClient){

	calculateMaxClient,err := client.CalculateMax(context.Background())

	if err != nil {
		log.Fatal("Error while making gRPC call to get max values. Error: ",err)
	}

	var dataPoints = []int32 {4,2,9,-9,88,-999,9999}

	ch := make(chan struct{})

	go func ()  {
		for _,data := range(dataPoints) {
			log.Printf("seding %d to max stream rpc call",data)
			err = calculateMaxClient.Send(&pb.MaxRequest{
				ReadingValue: data,
			})

			if err!= nil {
				log.Fatal("Error while sending max stream request. Error: ",err)
			}
		
			time.Sleep(1*time.Second)
		}
		
		err = calculateMaxClient.CloseSend()
		if err != nil {
			log.Fatal("Error while closing stream. Error: ",err)
		}
		
	}()

	go func ()  {
		defer func ()  {
			ch <- struct{}{}
		}()

		for {
			response,err := calculateMaxClient.Recv()

			if (err == io.EOF){
				break;
			}

			if (err != nil){
				log.Fatal("Error while receiving data from stream. Error: ",err)
			}

			log.Print("Response received via stream ",response.GetResult())
			
		}
	}()

	<- ch
}