package mysqlRepository

import (
	"feedProject/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(feed *models.Feed) error {
	return r.db.Create(feed).Error
}

func (r *Repository) Update(feed *models.Feed) error {
	return r.db.Save(feed).Error
}

func (r *Repository) DeleteByID(id int) error {
	return r.db.Delete(&models.Feed{}, id).Error
}

func (r *Repository) FindByID(id int) (*models.Feed, error) {
	var feed models.Feed
	err := r.db.First(&feed, id).Error
	if err != nil {
		return nil, err
	}
	return &feed, nil
}

// 根据publishID获取feeds
func (r *Repository) FindFeedsByPusherID(publisherID int, page int, pageSize int) ([]models.Feed, error) {
	var feeds []models.Feed
	offset := (page - 1) * pageSize
	err := r.db.Model(&models.Feed{}).
		Where("publisher_id = ?", publisherID).
		Order("publish_time DESC"). // 按照 PublishTime 降序排序
		Limit(pageSize).
		Offset(offset).
		Find(&feeds).Error
	return feeds, err
}
