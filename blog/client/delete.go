package main

import (
	"context"
	pb "grpcLearning/blog/proto"
	"log"
)

func deleteBlog(grpcClient pb.BlogServiceClient, id string){

	log.Printf("Making rpc call to delete a blog with id: %s\n",id)

	_,err := grpcClient.DeleteBlog(context.Background(), &pb.BlogId{
		Id: id,
	})

	if err != nil {
		log.Printf("Error while deleting the blog. Error: %v",err)
		return
	}

	log.Println("successfully deleted the blog.")
}