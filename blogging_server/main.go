// Package main implements a server for Bloggig service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/wrappers"
	bloggingapi "go.blogging.api/bloggingapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"math/rand"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	globalStore = make(map[uint64]*bloggingapi.BlogPost)
)

// server is used to implement Blogging API server.
type server struct {
	bloggingapi.UnimplementedBloggingapiServer
}

func (s *server) CreateBlog(ctx context.Context, blogPost *bloggingapi.BlogPost) (*bloggingapi.BlogPostWithUid, error) {
	blogPostId := rand.Uint64()
	log.Printf("Created Blog post: %v", blogPostId)
	globalStore[blogPostId] = blogPost
	return &bloggingapi.BlogPostWithUid{PostID: blogPostId, Post: blogPost}, nil
}

func (s *server) ReadBlog(ctx context.Context, in *bloggingapi.BlogRequest) (*bloggingapi.BlogPostWithUid, error) {
	blogPostId := in.GetPostid()
	log.Printf("Reading Blog post: %v", blogPostId)
	blogPost := globalStore[blogPostId]
	return &bloggingapi.BlogPostWithUid{PostID: blogPostId, Post: blogPost}, nil
}

func (s *server) DeleteBlog(ctx context.Context, in *bloggingapi.BlogRequest) (*wrappers.StringValue, error) {
	blogPostId := in.GetPostid()
	_, exists := globalStore[blogPostId]

	if exists {
		delete(globalStore, blogPostId)
		response := &wrappers.StringValue{
			Value: fmt.Sprintf("%s %d %s", "Blog post ", blogPostId, " deleted successfully"),
		}
		return response, nil
	}

	failResponse := &wrappers.StringValue{
		Value: fmt.Sprintf("%s %d %s", "Blog post ", blogPostId, " does not exist"),
	}
	return failResponse, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bloggingapi.RegisterBloggingapiServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
