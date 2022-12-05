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
	"google.golang.org/protobuf/types/known/emptypb"
)
func (s *Server) UpdateBlog(ctx context.Context, request *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("received request to update blog %v\n",request)

	objectId,err := primitive.ObjectIDFromHex(request.Id)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Couldn't parse the object id. Error %v\n",err))
	}

	updatedBlog := &BlogItem{
		AuthorId: request.AuthorId,
		Title: request.Title,
		Content: request.Content,
	}

	//if we are not passing some feilds here for that document, then it setting it to empty string for that feild
	response ,err := collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set":updatedBlog} )

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while updating the blog. Error: %s\n ",err))
	}

	if response.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Blog not found")
	}

	log.Print("Successfully updated the blog.")
	return &emptypb.Empty{},nil
}