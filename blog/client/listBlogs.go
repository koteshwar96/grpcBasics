package main

import (
	"context"
	pb "grpcLearning/blog/proto"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(client pb.BlogServiceClient) {

	log.Printf("Making a rpc call to get all the blogs.")

	listClient, err := client.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Printf("Error while making the rpc call to get all the blogs. Error: %v",err)
	}


	for {
		
		blog,err := listClient.Recv()

		if err != nil {
			if err == io.EOF {
				return
			}

			log.Printf("Error while receiving the blog stream from server. Error %v",err)
			return	
		}

		log.Println(blog)
	}
}