package main

import (
	pb "grpcLearning/calculator/proto"
	"io"
	"log"
	"math"
)

func (s *Server) CalculateMax( stream pb.CalculatorService_CalculateMaxServer) error {

	log.Print("Received calculateMax call to the server")

	var currentMaximum int32 = math.MinInt32

	for {
		maxRequest, err := stream.Recv()

		if (err == io.EOF){
			return nil
		}

		if (err != nil){
			log.Fatal("Error while receiving data from the stream of calculate max. Error: ",err)
		}

		maxRequest.GetReadingValue()
		currentMaximum = int32 (math.Max(float64(currentMaximum), float64(maxRequest.GetReadingValue())))

		err = stream.Send(&pb.MaxResponse{
			Result: currentMaximum,
		})

		if err != nil {
			log.Fatal("Error while sending max stream. Error: ",err)
		}
		
	}

}