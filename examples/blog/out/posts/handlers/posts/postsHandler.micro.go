package posts

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	server "github.com/micro/go-micro/v2/server"
	proto_posts "github.com/v1990/protoc-gen-goku/examples/blog/proto_go/posts"
)

type PostsHandler interface {
	Query(ctx context.Context, in *proto_posts.QueryRequest, out *proto_posts.QueryResponse) error

	Save(ctx context.Context, in *proto_posts.SaveRequest, out *proto_posts.SaveResponse) error

	Delete(ctx context.Context, in *proto_posts.DeleteRequest, out *proto_posts.DeleteResponse) error

	Report(ctx context.Context, in *proto_posts.ReportRequest, out *empty.Empty) error
}

func RegisterPostsHandler(s server.Server, opts ...server.HandlerOption) error {
	return proto_posts.RegisterPostsHandler(s, newPostsHandler(), opts...)
}

func newPostsHandler() *postsHandler {
	return new(postsHandler)
}

type postsProxyHandler struct {
	H *postsHandler
}

func (t *postsProxyHandler) Query(ctx context.Context, in *proto_posts.QueryRequest, out *proto_posts.QueryResponse) error {
	return t.H.Query(ctx, in, out)
}

func (t *postsProxyHandler) Save(ctx context.Context, in *proto_posts.SaveRequest, out *proto_posts.SaveResponse) error {
	return t.H.Save(ctx, in, out)
}

func (t *postsProxyHandler) Delete(ctx context.Context, in *proto_posts.DeleteRequest, out *proto_posts.DeleteResponse) error {
	return t.H.Delete(ctx, in, out)
}

func (t *postsProxyHandler) Report(ctx context.Context, in *proto_posts.ReportRequest, out *empty.Empty) error {
	return t.H.Report(ctx, in, out)
}
