package main

import (
	"context"
	"fmt"
	pb "grpcLearning/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, request *pb.Blog) (*pb.BlogId, error) {

	log.Printf("Create blog was called with %v \n",request)

	// we are not passing id field here because, mongo db will create it for us and return it
	blog := BlogItem {
		AuthorId: request.AuthorId,
		Title: request.Title,
		Content: request.Content,
	}

	response,err := collection.InsertOne(ctx, blog)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to insert blog into the DB. Error: %v\n",err))
	}

	id,ok := response.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, "Couldn't convert into id")
	}

	return &pb.BlogId{
		Id: id.Hex(),
	},nil
}