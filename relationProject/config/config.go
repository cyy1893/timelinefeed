package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "root:qqabc123@tcp(192.168.161.213:30922)/my_relation_schema?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                    dsn,  // DSN data source name
		DefaultStringSize:      191,  // string 类型字段的默认长度
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	}), &gorm.Config{SkipDefaultTransaction: false})

	return db, err
}
