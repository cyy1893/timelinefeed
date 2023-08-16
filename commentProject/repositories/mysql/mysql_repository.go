package mysqlRepository

import (
	"commentProject/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateComment 创建 Comment 记录
func (r *Repository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

// UpdateComment 更新 Comment 记录
func (r *Repository) UpdateComment(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

// DeleteComment 删除 Comment 记录
func (r *Repository) DeleteComment(comment *models.Comment) error {
	return r.db.Delete(comment).Error
}

// GetCommentByID 根据 CommentID 获取 Comment 记录
func (r *Repository) GetCommentByID(commentID int) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.First(&comment, commentID).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// GetCommentCountByFeedID 根据 FeedID 获取 CommentCount 记录的数量
func (r *Repository) GetCommentCountByFeedID(feedID int) (int, error) {
	var count int64
	err := r.db.Model(&models.Comment{}).Where("feed_id = ?", feedID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

// GetCommentsByUserID 根据UserID 获取 Comments
func (r *Repository) GetCommentsByUserID(userID int) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Model(&models.Comment{}).Where("user_id = ?", userID).Find(comments).Error
	return comments, err
}
