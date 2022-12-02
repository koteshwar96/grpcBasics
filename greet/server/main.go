package main

import (
	pb "grpcLearning/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct{
	pb.GreetServiceServer
}

func main(){

	lis,err := net.Listen("tcp",addr)

	if err != nil {
		log.Fatalf("Unable to start the server on the address %s. Error: %s",addr,err)
	}

	log.Printf("server started and listening on %s",lis.Addr())

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &Server{})
	
	if err = grpcServer.Serve(lis); err!=nil {
		log.Fatalf("Failed to serve grpc. Error : %s",err)
	}

}