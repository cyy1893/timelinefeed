package models

import (
	"gorm.io/gorm"
	"time"
)

type Feed struct {
	gorm.Model
	Content     string    `gorm:"size:255" json:"content"`
	PublisherID int       `json:"publisher_id"`
	PublishTime time.Time `json:"publish_time"`
}
