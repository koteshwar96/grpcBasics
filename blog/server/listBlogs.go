package main

import (
	"context"
	"fmt"
	pb "grpcLearning/blog/proto"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)
func (s *Server) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {

	log.Printf("Received request to list all the blogs\n")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error while getting all the blogs, Error: %v",err))
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		blog := &BlogItem{}

		err = cur.Decode(blog)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding the blog. Error: %v",err))
		}

		stream.Send(documentToBlogProto(blog))

	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error while getting the blogs. Error: %v",err))
	}

	return nil
}