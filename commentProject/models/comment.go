package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID         int    `json:"user_id"`
	FeedID         int    `json:"feed_id"`
	CommentContent string `json:"comment_content" gorm:"size:255"`
}

type CommentCount struct {
	FeedID       int
	CommentCount int
}
