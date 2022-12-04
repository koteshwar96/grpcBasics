package main


import (
	pb "grpcLearning/blog/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var host string = "localhost:50052"
func main(){

	// this variable should be set accordingly if tls is needed or not
	tls := false

	options := []grpc.DialOption{}

	if tls {
		// ca.crt is the certificate of the certificate authority which is used to validate the server cert
		caCrtFile := "ssl/ca.crt";

		creds,err := credentials.NewClientTLSFromFile(caCrtFile, "")
		if err != nil {
			log.Fatal("Error while creating credentials for grpc SSL. Error: ",err)
		}
		options = append(options, grpc.WithTransportCredentials(creds))
	} else {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials() ))
	}

	connection,err := grpc.Dial(host, options...)

	if err != nil {
		log.Fatalf("Error while dailing in the grpc client. Error: %s",err)
	}

	defer connection.Close()

	blogClient := pb.NewBlogServiceClient(connection)

	blogId := createBlog(blogClient)

	readBlog(blogClient, blogId)
	// Lets now call with some random string as uuid which will throw error
	readBlog(blogClient, "randomUnknownUUId")

}