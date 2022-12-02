package main

import (
	pb "grpcLearning/calculator/proto"
	"io"
	"log"
)

func (s *Server) CalculateAverage(stream pb.CalculatorService_CalculateAverageServer) error {

	log.Print("Received a gRPC call to determine average..")

	var readings []int32

	for {
		averageRequest, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error while receiving the data. Error: ", err)
		}

		log.Print("Received the next reading: ", averageRequest)
		readings = append(readings, averageRequest.GetReadingValue())
	}

	// one could trigger this function inside the loop as soon it encounters EOF
	return calculate(readings, stream)
}

func calculate(readings []int32, stream pb.CalculatorService_CalculateAverageServer) error {
	var sum int32
	for _, value := range readings {
		sum += value
	}

	return stream.SendAndClose(&pb.AverageResponse{
		ReadingAverage: float64 (sum) /float64 (len(readings)),
	})
}
