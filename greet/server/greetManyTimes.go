package main

import (
	"fmt"
	pb "grpcLearning/greet/proto"
	"log"
)
func (s *Server) GreetManyTimes(request *pb.GreetRequest, response  pb.GreetService_GreetManyTimesServer) error {
	log.Print("Received a request to greet many times.")

	for i:=1 ; i<=10 ; i++ {
		response.Send(&pb.GreetResponse{
			Response: fmt.Sprintf("Hello %s, time : %d",request.GetName(),i),
		})
	}
	
	return nil
}