package follow_server

import (
	"context"
	"relationProject/grpc/follow"
	"relationProject/services"
)

type FollowServiceServerImpl struct {
	follow.FollowServiceServer
	followService *services.FollowService
}

func NewFollowServiceServerImpl(followService *services.FollowService) *FollowServiceServerImpl {
	return &FollowServiceServerImpl{
		followService: followService,
	}
}

func (f *FollowServiceServerImpl) GetFollowersByFollowingID(ctx context.Context, request *follow.FollowingIDRequest) (*follow.FollowersResponse, error) {
	followers, err := f.followService.GetFollowersByFollowingID(uint(request.FollowingId))
	if err != nil {
		return nil, err
	}

	// 将 `[]uint` 转换为 `[]uint32` 类型
	convertedFollowers := make([]uint32, len(followers))
	for i, follower := range followers {
		convertedFollowers[i] = uint32(follower)
	}

	return &follow.FollowersResponse{
		FollowerIds: convertedFollowers,
	}, nil
}

func (f *FollowServiceServerImpl) mustEmbedUnimplementedFollowServiceServer() {}
