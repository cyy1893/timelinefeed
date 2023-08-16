package comment_server

import (
	"commentProject/grpc/comment"
	"commentProject/services"
	"context"
)

type CommentServiceServerImpl struct {
	comment.CommentServiceServer
	commentService *services.CommentService
}

func NewCommentServiceGrpcImpl(commentService *services.CommentService) *CommentServiceServerImpl {
	return &CommentServiceServerImpl{
		commentService: commentService,
	}
}

func (c *CommentServiceServerImpl) GetCommentCountByFeedID(ctx context.Context, req *comment.CommentCountRequest) (*comment.CommentCountResponse, error) {
	commentCount, err := c.commentService.GetCommentCountByFeedID(ctx, int(req.FeedId))
	if err != nil {
		return nil, err
	}

	return &comment.CommentCountResponse{
		Count: int32(commentCount),
	}, nil

}

//func (c *CommentServiceServerImpl) GetCommentsByUserID(ctx context.Context, req *comment.CommentsByUserIDRequest) (*comment.CommentsByUserIDResponse, error) {
//
//}

func (c *CommentServiceServerImpl) mustEmbedUnimplementedCommentServiceServer() {}
