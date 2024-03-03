/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
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

// server is used to implement helloworld.GreeterServer.
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
