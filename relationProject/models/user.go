package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255"`
	Avatar   string `gorm:"size:255"`
	// 其他字段
}
