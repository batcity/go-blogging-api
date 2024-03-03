// Package main implements a server for Bloggig service.
package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/wrappers"
	bloggingapi "go.blogging.api/bloggingapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port                     = flag.Int("port", 50051, "The server port")
	globalStore              = make(map[uint32]*bloggingapi.BlogPost)
	globallyUniqueUid uint32 = 0
)

// server is used to implement Blogging API server.
type server struct {
	bloggingapi.UnimplementedBloggingapiServer
}

func (s *server) CreateBlog(ctx context.Context, blogPost *bloggingapi.BlogPost) (*bloggingapi.BlogPostWithUid, error) {
	blogPostId := globallyUniqueUid
	log.Printf("Created Blog post: %v", blogPostId)
	globalStore[blogPostId] = blogPost
	globallyUniqueUid++
	return &bloggingapi.BlogPostWithUid{PostID: blogPostId, Post: blogPost}, nil
}

func (s *server) ReadBlog(ctx context.Context, in *bloggingapi.BlogPostID) (*bloggingapi.BlogPostWithUid, error) {
	blogPostId := in.GetPostID()
	log.Printf("Reading Blog post: %v", blogPostId)
	blogPost := globalStore[blogPostId]
	return &bloggingapi.BlogPostWithUid{PostID: blogPostId, Post: blogPost}, nil
}

func (s *server) UpdateBlog(ctx context.Context, blogPostWithUid *bloggingapi.BlogPostWithUid) (*bloggingapi.BlogPostWithUid, error) {
	blogPostId := blogPostWithUid.GetPostID()
	_, exists := globalStore[blogPostId]

	if exists {
		log.Printf("Updating Blog post: %v", blogPostId)
		globalStore[blogPostId] = blogPostWithUid.GetPost()
		return &bloggingapi.BlogPostWithUid{PostID: blogPostId, Post: globalStore[blogPostId]}, nil
	}

	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%s %d %s", "Blog post", blogPostId, "does not exist"))
}

func (s *server) DeleteBlog(ctx context.Context, in *bloggingapi.BlogPostID) (*wrappers.StringValue, error) {
	blogPostId := in.GetPostID()
	_, exists := globalStore[blogPostId]

	if exists {
		log.Printf("Deleting Blog post: %v", blogPostId)
		delete(globalStore, blogPostId)
		response := &wrappers.StringValue{
			Value: fmt.Sprintf("%s %d %s", "Blog post", blogPostId, "deleted successfully"),
		}
		return response, nil
	}

	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%s %d %s", "Blog post", blogPostId, "does not exist"))
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
