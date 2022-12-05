package main

import (
	"context"
	pb "grpcLearning/blog/proto"
	"log"
)
func updateBlog(client pb.BlogServiceClient, id string){
	log.Printf("Making call to update the blog with blog id : %s",id)

	newBlog := &pb.Blog{
		Id: id,
		Content: "Content is updated",
		Title: "new Title",
	}

	_,err := client.UpdateBlog(context.Background(), newBlog)

	// One of the key things to remember when ever an error occurs in golang, we can either do a fatal and exit from the program or handle gracefully. 
	// Make sure we are executing the after this block, check if we should execute it or not
	if err!= nil {
		log.Printf("Error while updating the blog. Error: %v\n",err)
		return
	}

	log.Printf("Successfully updated the blog with blog id : %s",id)

}
