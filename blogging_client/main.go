// Package main implements a client for Blogging API service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	bloggingapi "go.blogging.api/bloggingapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr        = flag.String("addr", "localhost:50051", "the address to connect to")
	sampleBlog  = bloggingapi.BlogPost{Title: "sample title", Content: "sample content"}
	updatedBlog = bloggingapi.BlogPost{Title: "updated sample title", Content: "updated sample content"}
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := bloggingapi.NewBloggingapiClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Creating a blog post
	blogPostWithUid, err := c.CreateBlog(ctx, &sampleBlog)
	if err != nil {
		log.Fatalf("could not create blog: %v", err)
	}
	log.Printf("Blog post created: %v", blogPostWithUid)
	var blogPostId = blogPostWithUid.GetPostID()

	// Reading a blog post
	blogPost, err := c.ReadBlog(ctx, &bloggingapi.BlogPostID{PostID: blogPostId})
	if err != nil {
		log.Fatalf("could not read blog: %v", err)
	}
	log.Printf("Retrieving post: %v", blogPost)

	// Updating a blog post
	var updatedBlogWithUid = bloggingapi.BlogPostWithUid{PostID: blogPostId, Post: &updatedBlog}
	updatedBlogPostWithUid, err := c.UpdateBlog(ctx, &updatedBlogWithUid)
	if err != nil {
		log.Fatalf("could not update blog: %v", err)
	}
	log.Printf("Blog post updated: %v", updatedBlogPostWithUid)

	// Deleting a blog post
	response, err := c.DeleteBlog(ctx, &bloggingapi.BlogPostID{PostID: blogPostId})
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%s", response)
}
