package main

import (
	"context"
	"log"

	pb "github.com/mnbedel/blog-poster/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("***createBlog was invoked***")

	blog := &pb.Blog{
		AuthorId: "Mustafa",
		Title:    "First blog by Mustafa",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been posted: %s\n", res.Id)

	return res.Id
}
