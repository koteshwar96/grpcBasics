package main

import (
	pb "grpcLearning/calculator/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var host string = "localhost:50052"
func main(){

	// this variable should be set to false if tls is not needed
	tls := true

	options := []grpc.DialOption{}

	if tls {
		// ca.crt is the certificate of the certificate authority which is used to validate the server cert
		caCrtFile := "ssl/ca.crt";

		creds,err := credentials.NewClientTLSFromFile(caCrtFile, "")
		if err != nil {
			log.Fatal("Error while creating credentials for grpc SSL. Error: ",err)
		}
		options = append(options, grpc.WithTransportCredentials(creds))
	} else {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials() ))
	}

	connection,err := grpc.Dial(host, options...)

	if err != nil {
		log.Fatalf("Error while dailing in the grpc client. Error: %s",err)
	}

	defer connection.Close()

	grpcClient := pb.NewCalculatorServiceClient(connection)

	// log.Println("\nMaking calculate sum gRPC call (unary)")
	calculateSum(grpcClient)

	// log.Println("\nMaking determinePrimeFactors (server stream)")
	// determinePrimeFactors(grpcClient)

	// log.Println("\nMaking calculateAverage gRPC call (client stream)")
	// calculateAverage(grpcClient)

	// log.Println("\nMaking MaxRPC Call (bi-directional stream)")
	// makeMaxRPCCall(grpcClient)

	// We have handled errors gracefully for this rpc call
	// log.Println("\nMaking sqrt Call")
	// calculateSqRt(grpcClient,64)
	// calculateSqRt(grpcClient,-64)

	// we are making grpc calls with deadline using context
	// we will get the result here as deadline is not crossed
	calculateSumWithDeadline(grpcClient, 7*time.Second)
	// we will get dead line exceeded here
	// calculateSumWithDeadline(grpcClient, 3*time.Second)

}