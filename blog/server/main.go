package main

import (
	"context"
	"log"
	"net"

	pb "grpcLearning/blog/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address string = "0.0.0.0:50052"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	// creating a client connection to the MongoDB
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal("Error while creating client connection to th MongoDB. Error: ",err)
	}

	client.Connect(context.Background())
	if err != nil {
		log.Fatal("Error while creating client connection to th MongoDB. Error: ",err)
	}

	collection = client.Database("blogdb").Collection("blog");


	// creating a server on the given address
	listener,err := net.Listen("tcp",address)
	if err != nil {
		log.Fatalf("Failed to listed on the address %s. Error: %s",address,err)
	}

	log.Print("Server starting on ",address)

	// set this variabel accordingly if tls is required or not
	tls := false
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

	pb.RegisterBlogServiceServer(grpcServer, &Server{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start grpc server on the listener. Error: %s",err)
	}

}