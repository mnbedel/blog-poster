package main

import (
	"log"

	pb "github.com/mnbedel/blog-poster/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)

	readBlog(c, id) // exist in mongo
	//readBlog(c, "invalid id") // not exist in mongo
	updateBlog(c, id) // update the last valid record
	listBlog(c)       // list all records
	deleteBlog(c, id) // delete the blog with the Id of id
}
