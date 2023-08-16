package config

import (
	"feedProject/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	DSN                    string
	DefaultStringSize      int
	DontSupportRenameIndex bool
}

var MySQLConfigInstance = MySQLConfig{
	DSN:                    "root:qqabc123@tcp(192.168.161.213:30922)/my_feed_schema?charset=utf8mb4&parseTime=True&loc=Local",
	DefaultStringSize:      191,
	DontSupportRenameIndex: true,
}

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                    MySQLConfigInstance.DSN,
		DefaultStringSize:      uint(MySQLConfigInstance.DefaultStringSize),
		DontSupportRenameIndex: MySQLConfigInstance.DontSupportRenameIndex,
	}), &gorm.Config{SkipDefaultTransaction: false})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Feed{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
