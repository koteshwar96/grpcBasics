package main

import (
	"log"
	"net"

	pb "grpcLearning/calculator/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	listener,err := net.Listen("tcp",address)
	if err != nil {
		log.Fatalf("Failed to listed on the address %s. Error: %s",address,err)
	}

	log.Print("Server starting on ",address)

	// set this to false if tls is not required
	tls := true
	var options = []grpc.ServerOption{}

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		creds,err := credentials.NewServerTLSFromFile(certFile,keyFile)
		if err != nil {
			log.Fatal("Error while creating tls crendentials from the certs. Error: ",err)
		}

		options = append(options, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(options...)
	
	defer grpcServer.Stop()

	pb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start grpc server on the listener. Error: %s",err)
	}

	
}