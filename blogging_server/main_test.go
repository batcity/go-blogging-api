package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	bloggingapi "go.blogging.api/bloggingapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var (
	lis               *bufconn.Listener
	ctx               = context.Background()
	sampleBlog        = bloggingapi.BlogPost{Title: "sample title", Content: "sample content"}
	sampleBlogWithUid = bloggingapi.BlogPostWithUid{PostID: 0, Post: &sampleBlog}
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	bloggingapi.RegisterBloggingapiServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func newTestClient(t *testing.T) bloggingapi.BloggingapiClient {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	t.Cleanup(func() { conn.Close() }) // Use t.Cleanup for proper cleanup
	return bloggingapi.NewBloggingapiClient(conn)
}

func TestCreateBlog(t *testing.T) {

	client := newTestClient(t)
	got, err := client.CreateBlog(ctx, &sampleBlog)

	if err != nil {
		t.Fatalf("CreateBlog failed: %v", err)
	}
	log.Printf("Response: %+v", got)

	want := &sampleBlogWithUid

	if got.String() != want.String() {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestReadBlog(t *testing.T) {

	client := newTestClient(t)
	got, err := client.ReadBlog(ctx, &bloggingapi.BlogPostID{PostID: 0})

	if err != nil {
		t.Fatalf("ReadBlog failed: %v", err)
	}
	log.Printf("Response: %+v", got)

	want := &sampleBlogWithUid
	if got.String() != want.String() {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestReadBlogErrorPath(t *testing.T) {

	client := newTestClient(t)
	_, err := client.ReadBlog(ctx, &bloggingapi.BlogPostID{PostID: 100})

	log.Printf("Response: %+v", err)
	got := err.Error()
	want := "rpc error: code = NotFound desc = Blog post 100 does not exist"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestUpdateBlog(t *testing.T) {

	client := newTestClient(t)
	updatedBlog := bloggingapi.BlogPost{Title: "updated sample title", Content: "updated sample content"}
	updatedBlogWithUid := bloggingapi.BlogPostWithUid{PostID: 0, Post: &updatedBlog}
	_, err := client.UpdateBlog(ctx, &updatedBlogWithUid)
	got, err := client.ReadBlog(ctx, &bloggingapi.BlogPostID{PostID: 0})

	if err != nil {
		t.Fatalf("UpdateBlog failed: %v", err)
	}
	log.Printf("Response: %+v", got)

	want := &updatedBlogWithUid
	if got.String() != want.String() {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestUpdateBlogErrorPath(t *testing.T) {

	client := newTestClient(t)
	updatedBlog := bloggingapi.BlogPost{Title: "updated sample title", Content: "updated sample content"}
	updatedBlogWithUid := bloggingapi.BlogPostWithUid{PostID: 100, Post: &updatedBlog}
	_, err := client.UpdateBlog(ctx, &updatedBlogWithUid)

	log.Printf("Response: %+v", err)
	got := err.Error()
	want := "rpc error: code = NotFound desc = Blog post 100 does not exist"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestDeleteBlog(t *testing.T) {

	client := newTestClient(t)
	got, err := client.DeleteBlog(ctx, &bloggingapi.BlogPostID{PostID: 0})

	if err != nil {
		t.Fatalf("DeleteBlog failed: %v", err)
	}
	log.Printf("Response: %+v", got)

	want := &wrappers.StringValue{Value: "Blog post 0 deleted successfully"}
	if got.String() != want.String() {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestDeleteBlogErrorPath(t *testing.T) {

	client := newTestClient(t)
	_, err := client.DeleteBlog(ctx, &bloggingapi.BlogPostID{PostID: 100})

	log.Printf("Response: %+v", err)
	got := err.Error()
	want := "rpc error: code = NotFound desc = Blog post 100 does not exist"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
