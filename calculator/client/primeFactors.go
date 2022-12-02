package main

import (
	"context"
	pb "grpcLearning/calculator/proto"
	"io"
	"log"
)
func determinePrimeFactors(grpcClient pb.CalculatorServiceClient){
	log.Print("Now making grpc call to get prime factors..")

	var number int32 = 975
	primeFactorsClient, err := grpcClient.DeterminePrimeFactors(context.Background(), &pb.PrimeFactorsRequest{
		Number: number,
	})

	if(err != nil){
		log.Fatalf("Error while making rpc call to get prime factors. Error: %s",err)
	}

	
	log.Printf("Prime factors of %d are, ",number)

	var primeFactors []int32

	for {
		primeFactorResponse, err := primeFactorsClient.Recv()

		if (err != nil){
			if (err == io.EOF){
				break;
			}
			log.Fatalf("Error while receiving prime factors. Error: %s ",err)
		}
		log.Printf("Next Prime factor is: (stream) %d",primeFactorResponse.GetPrimeFactors())

		primeFactors = append(primeFactors, primeFactorResponse.GetPrimeFactors())
	}
	log.Printf("Prime factors of %d are: %d",number,primeFactors)
}