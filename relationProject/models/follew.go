package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowerID  uint
	FollowingID uint
}
