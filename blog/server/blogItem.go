package main

import ("go.mongodb.org/mongo-driver/bson/primitive"
	pb "grpcLearning/blog/proto"
)

// All field names must start with caps to be marshal and unmarshal. We are creating this struct for the mongo driver to be able read and write from db
type BlogItem struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string `bson:"author_id"`
	Title string	`bson:"title"`
	Content string  `bson:"content"`
}

func documentToBlogProto (blog *BlogItem) (*pb.Blog){
	return &pb.Blog{
		Id: blog.ID.Hex(),
		AuthorId: blog.AuthorId,
		Title: blog.Title,
		Content: blog.Content,
	}
}
