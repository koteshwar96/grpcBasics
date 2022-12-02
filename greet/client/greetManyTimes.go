package main

import (
	"context"
	pb "grpcLearning/greet/proto"
	"io"
	"log"
)
func greetManyTimes(client pb.GreetServiceClient ) {
	log.Print("Making multiple greet request call.")

	greetManyTimesClient, err := client.GreetManyTimes(context.Background(), &pb.GreetRequest{
		Name: "Koteshwar",
	})

	if err!= nil {
	
		log.Fatalf("Error while making greet many times grpc call. Error: %s",err)
	}

	for {
		message, err2 := greetManyTimesClient.Recv()

		if err2 != nil {
			if err2 == io.EOF {
				break;
			}
			log.Fatalf("Error while receiving the message. Error: %s",err2)
		}

		log.Print("Message Received is: ",message.GetResponse())
	}
}