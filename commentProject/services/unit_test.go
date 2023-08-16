package services

//// 示例函数，用于创建 CommentCountRedisRepository 实例
//func createCommentCountRepository() *redisRepository.Repository {
//	// 建立 Redis 连接
//	redisClient := redis.NewClient(&redis.Options{
//		Addr:     "192.168.161.213:32600", // Redis 服务器的地址和端口
//		Password: "",                      // Redis 服务器的密码，如果没有密码则为空字符串
//		DB:       0,                       // Redis 数据库的索引，默认为 0
//	})
//
//	// 创建 CommentCountRedisRepository 实例
//	commentCountRepo := redisRepository.NewRepository(redisClient)
//
//	return commentCountRepo
//}
//
//func createCommentCountRepositoryMysql() *mysqlRepository.Repository {
//	dsn := "root:qqabc123@tcp(192.168.161.213:30922)/my_comment_schema?charset=utf8mb4&parseTime=True&loc=Local"
//
//	db, err := gorm.Open(mysql.New(mysql.Config{
//		DSN:                    dsn,  // DSN data source name
//		DefaultStringSize:      191,  // string 类型字段的默认长度
//		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
//	}), &gorm.Config{SkipDefaultTransaction: false})
//
//	err = db.AutoMigrate(&models.Comment{}, &models.CommentCount{})
//	if err != nil {
//		panic(err)
//	}
//
//	return mysqlRepository.NewCommentCountRepository(db)
//}
//
//func createCommentCountRepositoryKafka() *kafkaRepository.Repository {
//	brokers := []string{"192.168.161.213:9092"}
//	topic := "test-topic"
//
//	// 创建 Kafka 生产者
//	producer, err := kafkaRepository.NewKafkaProducer(brokers, topic)
//	if err != nil {
//		panic(err)
//	}
//
//	return kafkaRepository.NewKafkaRepository(producer)
//
//}
//
//func newCommentCountService() *CommentCountService {
//
//	// 创建测试数据和依赖项
//	redisRepo := createCommentCountRepository()      // 替换为实际的 Redis 存储库
//	mysqlRepo := createCommentCountRepositoryMysql() // 替换为实际的 MySQL 存储库
//	kafkaRepo := createCommentCountRepositoryKafka()
//	commentCountService := NewCommentCountService(mysqlRepo, redisRepo, kafkaRepo)
//	return commentCountService
//}
//
//func TestCommentCountService(t *testing.T) {
//	service := newCommentCountService()
//
//	commentCount, err := service.GetCommentCountByID(context.Background(), 312)
//	if err != nil {
//		panic(err)
//	}
//
//	commentCount.CommentCount = 1000
//	commentCount.FeedID = 1000
//
//	err = service.UpdateCommentCount(context.Background(), commentCount)
//
//}
//
////=== RUN   TestCommentCountService_KafkaIncrease
////--- PASS: TestCommentCountService_KafkaIncrease (13.74s)
////PASS
//
//func TestCommentCountService_KafkaIncrease(t *testing.T) {
//	service := newCommentCountService()
//
//	for i := 0; i < 10000; i++ {
//		err := service.IncreaseCommentCountByID(context.Background(), 312)
//		if err != nil {
//			panic(err)
//		}
//	}
//
//}
