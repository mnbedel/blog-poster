package main

import (
	"context"
	"log"

	pb "github.com/mnbedel/blog-poster/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("***updateBlog was invoked***")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Roger",
		Title:    "Go Go Go",
		Content:  "blog content update",
	}

	_, err := c.UpdateBlog(context.TODO(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated..!")
}
