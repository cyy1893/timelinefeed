package services

import (
	"relationProject/models"
	mysqlRepository "relationProject/repositories/mysql"
)

type UserService struct {
	repo *mysqlRepository.Repository
}

func NewUserService(repo *mysqlRepository.Repository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser 创建 User 记录
func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}

// UpdateUser 更新 User 记录
func (s *UserService) UpdateUser(user *models.User) error {
	return s.repo.UpdateUser(user)
}

// DeleteUserByID 根据 ID 删除 User 记录
func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}

// GetUserByID 根据 ID 获取 User 记录
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
