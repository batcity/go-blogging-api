// Package main implements a server for Bloggig service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
  "math/rand"
	pb "go.blogging.api/bloggingapi"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
  globalStore = inMemoryStore{data: make(map[uint64]string)}
)

// server is used to implement Blogging API server.
type server struct {
	pb.UnimplementedBloggingapiServer
}

// type BlogPost struct {
//   pb.BlogPost
// }

type inMemoryStore struct {
    data map[uint64]string
}

// func (s *inMemoryStore) Set(ctx context.Context, req *SetReq) (*SetResp, error) {
//     s.data[req.Key] = req.Value
//     return &SetResp{}, nil
// }

func (s *inMemoryStore) SaveBlog(ctx context.Context, in *pb.BlogPost) {
   id := in.GetPostID()
   s.data[id] = in.GetTitle()
   return
}

func (s *server) CreateBlog(ctx context.Context, in *pb.BlogPost) (*pb.BlogPost, error) {
  id := rand.Uint64()
  in.PostID = id
	log.Printf("Created Blog post: %v", id)
  globalStore.SaveBlog(ctx, in)
	return &pb.BlogPost{PostID: id, Title: in.GetTitle()}, nil
}

func (s *server) ReadBlog(ctx context.Context, in *pb.BlogRequest) (*pb.BlogPost, error) {
  id := in.GetPostid()
  log.Printf("Reading Blog post: %v", id)
  title := globalStore.data[id]
	return &pb.BlogPost{PostID: id, Title: title}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBloggingapiServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
