package main

import (
	"context"
	pb "grpcLearning/blog/proto"
	"log"
)

func createBlog(client pb.BlogServiceClient) string{
	log.Println("We are making create blog grpc call")

	blog := &pb.Blog{
		AuthorId: "Koteshwar",
		Title: "grpc Learning",
		Content: "This blog is about creation of basic grpc calls",

	}

	response,err := client.CreateBlog(context.Background(),blog)

	if err != nil {
		log.Fatal("Error while creating a blog. Error: %v\n",err) 
	}

	log.Printf("Successfully created blog and its id is %s",response.Id)
	return response.Id
}