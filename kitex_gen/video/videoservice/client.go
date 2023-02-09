// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	video "github.com/wen-flower/easy-douyin/kitex_gen/video"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateVideo(ctx context.Context, param *video.CreateVideoParam, callOptions ...callopt.Option) (r *video.CreateVideoResp, err error)
	PublishList(ctx context.Context, param *video.PublishListParam, callOptions ...callopt.Option) (r *video.PublishListResp, err error)
	VideoFeed(ctx context.Context, param *video.VideoFeedParam, callOptions ...callopt.Option) (r *video.VideoFeedResp, err error)
	FavoriteVideo(ctx context.Context, param *video.FavoriteVideoParam, callOptions ...callopt.Option) (r *video.FavoriteVideoResp, err error)
	FavoriteList(ctx context.Context, param *video.FavoriteListParam, callOptions ...callopt.Option) (r *video.FavoriteListResp, err error)
	CommentVideo(ctx context.Context, param *video.CommentVideoParam, callOptions ...callopt.Option) (r *video.CommentVideoResp, err error)
	DeleteComment(ctx context.Context, param *video.DeleteCommentParam, callOptions ...callopt.Option) (r *video.DeleteCommentResp, err error)
	CommentList(ctx context.Context, param *video.CommentListParam, callOptions ...callopt.Option) (r *video.CommentListResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) CreateVideo(ctx context.Context, param *video.CreateVideoParam, callOptions ...callopt.Option) (r *video.CreateVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVideo(ctx, param)
}

func (p *kVideoServiceClient) PublishList(ctx context.Context, param *video.PublishListParam, callOptions ...callopt.Option) (r *video.PublishListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, param)
}

func (p *kVideoServiceClient) VideoFeed(ctx context.Context, param *video.VideoFeedParam, callOptions ...callopt.Option) (r *video.VideoFeedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoFeed(ctx, param)
}

func (p *kVideoServiceClient) FavoriteVideo(ctx context.Context, param *video.FavoriteVideoParam, callOptions ...callopt.Option) (r *video.FavoriteVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteVideo(ctx, param)
}

func (p *kVideoServiceClient) FavoriteList(ctx context.Context, param *video.FavoriteListParam, callOptions ...callopt.Option) (r *video.FavoriteListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, param)
}

func (p *kVideoServiceClient) CommentVideo(ctx context.Context, param *video.CommentVideoParam, callOptions ...callopt.Option) (r *video.CommentVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentVideo(ctx, param)
}

func (p *kVideoServiceClient) DeleteComment(ctx context.Context, param *video.DeleteCommentParam, callOptions ...callopt.Option) (r *video.DeleteCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, param)
}

func (p *kVideoServiceClient) CommentList(ctx context.Context, param *video.CommentListParam, callOptions ...callopt.Option) (r *video.CommentListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, param)
}