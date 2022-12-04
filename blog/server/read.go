package main

import (
	"context"
	"fmt"
	pb "grpcLearning/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, request *pb.BlogId) (*pb.Blog, error) {

	log.Printf("Received request to get blog with id %v\n",request)

	objectID,err := primitive.ObjectIDFromHex(request.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Couldn;t parse id to object Id")
	}

	filter := bson.M{"_id":objectID}
	response := collection.FindOne(ctx,filter)

	var blog = &BlogItem{}
	err = response.Decode(blog)

	if err != nil {
		log.Printf("Error while decoding the blog. Error: %v\n",err)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Error while decoding the blog. Error: %v\n",err))
	}

	return documentToBlogProto(blog),nil

}
	