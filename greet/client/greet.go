package main

import (
	"context"
	pb "grpcLearning/greet/proto"
	"log"
)

func sendGreetings(client pb.GreetServiceClient){
	log.Println("Making RPC call to greet")

	response,err := client.Greet(context.Background(), &pb.GreetRequest{
		Name: "Koteshwar",
	})

	if err != nil {
		log.Fatalf("Failed to make an RPC call. Error: %s",err)
	}

	log.Printf("Reponse is : %s",response.GetResponse())

}