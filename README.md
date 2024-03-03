# go-blogging-api

Simple CRUD blogging app built using golang and grpc

## How to spin up the blogging server:

Run the following command:

```
go run blogging_server/main.go
```

## How to test the app:
### (Note: Make sure to install grpcurl first)

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
grpcurl -plaintext -format text -d 'postID: <postID>' localhost:50051 bloggingapi.bloggingapi.ReadBlog
```

3. You can Update a blog post using the following command:

```
grpcurl -plaintext -d '{
  "postID": "<post_id>",
  "post": {
    "title": "My First Blog Post updated",
    "content": "This is the content of my first blog post.",
    "author": "John Doe",
    "publicationdate": "2024-03-03T00:00:00Z",
    "tags": ["grpc", "protobuf"]
  }
}' localhost:50051 bloggingapi.bloggingapi.UpdateBlog
```

4. You can delete a blog post using the following command:

```
grpcurl -plaintext -format text -d 'postID: <postID>' localhost:50051 bloggingapi.bloggingapi.DeleteBlog
```
