package main

import (
	"commentProject/config"
	"commentProject/grpc/comment"
	comment_server "commentProject/grpc/comment_server"
	kafkaRepository "commentProject/repositories/kafka"
	redisRepository "commentProject/repositories/redis"
	"github.com/redis/go-redis/v9"

	"commentProject/handlers"
	"commentProject/models"
	mysqlRepository "commentProject/repositories/mysql"
	"commentProject/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	dbConfig := config.LoadDatabaseConfig()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                    dbConfig.DSN,
		DefaultStringSize:      uint(dbConfig.DefaultStringSize),
		DontSupportRenameIndex: dbConfig.DontSupportRenameIndex,
	}), &gorm.Config{SkipDefaultTransaction: false})

	err = db.AutoMigrate(&models.Comment{})
	if err != nil {
		panic(err)
	}

	rdbConfig := config.LoadRedisConfig()

	rdb := redis.NewClient(&redis.Options{
		Addr:     rdbConfig.Address,
		Password: rdbConfig.Password,
		DB:       rdbConfig.DB,
	})

	repository := mysqlRepository.NewRepository(db)
	rdbRepository := redisRepository.NewRepository(rdb)
	kfkRepository := kafkaRepository.NewKafkaRepository("192.168.161.213:9092", "comment")
	service := services.NewCommentService(repository, rdbRepository, kfkRepository)
	handler := handlers.NewCommentHandler(service)

	router := gin.Default()

	router.GET("/count_service/comment/:id", handler.GetCommentByID)
	router.POST("/count_service/comment", handler.CreateComment)
	router.PUT("/count_service/comment", handler.UpdateComment)
	router.DELETE("/count_service/comment/:id", handler.DeleteComment)
	router.GET("/count_service/count/feed/:id", handler.GetCommentCountByFeedID)

	// 启动 Gin 引擎
	go func() {
		err = router.Run(":8082")
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	// grpc server start
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server listening on port 8002")

	// 启动 gRPC 服务
	s := grpc.NewServer()
	comment.RegisterCommentServiceServer(s, comment_server.NewCommentServiceGrpcImpl(service))

	go func() {
		log.Println("gRPC server starting to serve...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	select {}
}
