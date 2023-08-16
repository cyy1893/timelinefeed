package mysqlRepository

import (
	"gorm.io/gorm"
	"relationProject/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// CreateFollow 创建 Follow 记录
func (r *Repository) CreateFollow(follow *models.Follow) error {
	return r.db.Create(follow).Error
}

// UpdateFollow 更新 Follow 记录
func (r *Repository) UpdateFollow(follow *models.Follow) error {
	return r.db.Save(follow).Error
}

// DeleteFollowByID 根据 ID 删除 Follow 记录
func (r *Repository) DeleteFollowByID(id uint) error {
	return r.db.Delete(&models.Follow{}, id).Error
}

// GetFollowByID 根据 ID 获取 Follow 记录
func (r *Repository) GetFollowByID(id uint) (*models.Follow, error) {
	var follow models.Follow
	err := r.db.First(&follow, id).Error
	if err != nil {
		return nil, err
	}
	return &follow, nil
}

// CreateUser 创建 User 记录
func (r *Repository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// UpdateUser 更新 User 记录
func (r *Repository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

// DeleteUserByID 根据 ID 删除 User 记录
func (r *Repository) DeleteUserByID(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// GetUserByID 根据 ID 获取 User 记录
func (r *Repository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 根据following_id获得所有的follower_id
func (r *Repository) GetFollowersByFollowingID(followingID uint) ([]uint, error) {
	var followers []models.Follow
	err := r.db.Where("following_id = ?", followingID).Find(&followers).Error
	if err != nil {
		return nil, err
	}
	var followerIDs []uint
	for _, follower := range followers {
		followerIDs = append(followerIDs, follower.FollowerID)
	}
	return followerIDs, nil
}
