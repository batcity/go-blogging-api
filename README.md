# go-blogging-api

Simple CRUD blogging app built using golang and grpc

## How to test the app:

1. You can create a blog post using the following command

```
  grpcurl -plaintext -d '{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "author": "John Doe",
  "publicationdate": "2024-03-03T00:00:00Z",
  "tags": ["grpc", "protobuf"]
}' localhost:50051 bloggingapi.bloggingapi.CreateBlog
```


2. You can retrieve a blog post using the following command:

```
grpcurl -plaintext -format text -d 'postid: 3134116674335633566' localhost:50051 bloggingapi.bloggingapi.ReadBlog
```
