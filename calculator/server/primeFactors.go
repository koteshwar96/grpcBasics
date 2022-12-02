package main

import (
	pb "grpcLearning/calculator/proto"
	"log"
)

func (s *Server) DeterminePrimeFactors(request *pb.PrimeFactorsRequest,stream pb.CalculatorService_DeterminePrimeFactorsServer) error {
	log.Printf("Received request to determine prime factors of %d",request.GetNumber())


	// I have moved the calculation of prime factors to another method. One can calculate here only and send as stream response then and there
	// it self when its exactly devisible instead of adding to an array and then sending from the array
	primeFactors := calculatePrimeFactors(request.GetNumber())

	for _, primeFactor := range primeFactors {
		stream.Send(&pb.PrimeFactorsResponse{
			PrimeFactors: primeFactor,
		})
	}
	return nil;
}

func calculatePrimeFactors(number int32) ([]int32 ) {
	primeFactors :=  []int32{};

	var divider int32 = 2

	for ;number > 1; {
		if number%divider == 0 {
			primeFactors = append(primeFactors, divider)
			number = number/divider
		} else {
			divider++
		}
		
	}

	return primeFactors
}