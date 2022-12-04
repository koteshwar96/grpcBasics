package main

import (
	"context"
	pb "grpcLearning/blog/proto"
	"log"
)
func readBlog (client pb.BlogServiceClient, id string ){
	log.Printf("making rpc call to get the blog details of %s",id)

	request := &pb.BlogId{
		Id: id,
	}

	response,err := client.ReadBlog(context.Background(),request)

	if err!= nil {
		log.Printf("Error while getting blog details. Error %v",err)
		return
	}

	log.Printf("Successfuly got the blog details. Blog details are, %v",response)
}