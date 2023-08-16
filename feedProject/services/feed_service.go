package services

import (
	"context"
	"feedProject/grpc/comment"
	"feedProject/grpc/follow"
	"feedProject/models"
	"feedProject/repositories/mysql"
	"google.golang.org/grpc"
)

type FeedService struct {
	feedRepo *mysqlRepository.Repository

	followServiceClient  follow.FollowServiceClient
	commentServiceClient comment.CommentServiceClient
}

func NewFeedService(feedRepo *mysqlRepository.Repository, followServiceConn grpc.ClientConnInterface, commentServiceConn grpc.ClientConnInterface) *FeedService {
	followServiceClient := follow.NewFollowServiceClient(followServiceConn)
	commentServiceClient := comment.NewCommentServiceClient(commentServiceConn)
	return &FeedService{
		feedRepo:             feedRepo,
		followServiceClient:  followServiceClient,
		commentServiceClient: commentServiceClient,
	}
}

func (s *FeedService) CreateFeed(feed *models.Feed) error {
	return s.feedRepo.Create(feed)
}

func (s *FeedService) UpdateFeed(feed *models.Feed) error {
	return s.feedRepo.Update(feed)
}

func (s *FeedService) DeleteFeedByID(id int) error {
	return s.feedRepo.DeleteByID(id)
}

func (s *FeedService) GetFeedByID(id int) (*models.Feed, error) {
	return s.feedRepo.FindByID(id)
}

// 根据publishID获取feeds
func (s *FeedService) GetFeedsByPusherID(pusherID int, page int, pageSize int) ([]models.Feed, error) {
	return s.feedRepo.FindFeedsByPusherID(pusherID, page, pageSize)
}

func (s *FeedService) GetFollowersByFollowingID(ctx context.Context, followingID *follow.FollowingIDRequest, opts ...grpc.CallOption) (*follow.FollowersResponse, error) {
	followers, err := s.followServiceClient.GetFollowersByFollowingID(ctx, followingID, opts...)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (s *FeedService) GetCommentCountByFeedID(ctx context.Context, feedID *comment.CommentCountRequest, opts ...grpc.CallOption) (*comment.CommentCountResponse, error) {
	commentCount, err := s.commentServiceClient.GetCommentCountByFeedID(ctx, feedID, opts...)
	if err != nil {
		return nil, err
	}
	return commentCount, nil
}
