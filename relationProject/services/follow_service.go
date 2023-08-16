package services

import (
	"relationProject/models"
	mysqlRepository "relationProject/repositories/mysql"
)

type FollowService struct {
	repo *mysqlRepository.Repository
}

func NewFollowService(repo *mysqlRepository.Repository) *FollowService {
	return &FollowService{repo: repo}
}

// CreateFollow 创建 Follow 记录
func (s *FollowService) CreateFollow(follow *models.Follow) error {
	return s.repo.CreateFollow(follow)
}

// UpdateFollow 更新 Follow 记录
func (s *FollowService) UpdateFollow(follow *models.Follow) error {
	return s.repo.UpdateFollow(follow)
}

// DeleteFollowByID 根据 ID 删除 Follow 记录
func (s *FollowService) DeleteFollowByID(id uint) error {
	return s.repo.DeleteFollowByID(id)
}

// GetFollowByID 根据 ID 获取 Follow 记录
func (s *FollowService) GetFollowByID(id uint) (*models.Follow, error) {
	return s.repo.GetFollowByID(id)
}

// GetFollowersByFollowingID 根据 FollowingID 获取 Followers
func (s *FollowService) GetFollowersByFollowingID(followingID uint) ([]uint, error) {
	return s.repo.GetFollowersByFollowingID(followingID)
}
