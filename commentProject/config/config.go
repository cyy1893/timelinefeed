package config

type DatabaseConfig struct {
	DSN                    string
	DefaultStringSize      int
	DontSupportRenameIndex bool
}

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func LoadDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DSN:                    "root:qqabc123@tcp(192.168.161.213:30922)/my_comment_schema?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:      191,
		DontSupportRenameIndex: true,
	}
}

func LoadRedisConfig() *RedisConfig {
	return &RedisConfig{
		Address:  "192.168.161.213:32600",
		Password: "",
		DB:       0,
	}
}
