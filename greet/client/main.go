package main

import (
	"log"

	pb "grpcLearning/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "localhost:50051"

func main() {

	connection, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to the grpc server. Error: %s", err)
	}

	defer connection.Close()

	log.Print("Successful connected to server.", connection.GetState().String())

	client := pb.NewGreetServiceClient(connection)

	// sendGreetings(client)
	greetManyTimes(client)

}
