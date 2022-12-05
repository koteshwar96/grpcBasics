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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, id *pb.BlogId) (*emptypb.Empty, error){

	log.Printf("Received request to delete the blog %v\n",id)

	objectId,err := primitive.ObjectIDFromHex(id.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Invalid objectId. Failed to convert to object id. Error: %v",err))
	}

	result,err := collection.DeleteOne(ctx, bson.M{"_id":objectId})

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while deleting the blog, Error: %v",err))
	}

	if result.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Blog Not found with the given id"))
	}

	return &emptypb.Empty{},nil
}