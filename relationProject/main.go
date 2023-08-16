package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
	"relationProject/config"
	"relationProject/grpc/follow"
	"relationProject/grpc/follow_server"
	"relationProject/handler"
	"relationProject/models"
	mysqlRepository "relationProject/repositories/mysql"
	"relationProject/services"
)

func main() {
	db, err := config.NewDB()

	err = db.AutoMigrate(&models.Follow{}, &models.User{})
	if err != nil {
		panic(err)
	}

	// 创建 UserRepository 实例
	repo := mysqlRepository.NewRepository(db)
	// 创建 UserService 实例
	userService := services.NewUserService(repo)
	// 创建 UserHandler 实例
	userHandler := handler.NewUserHandler(userService)
	// 创建 FollowService 实例
	followService := services.NewFollowService(repo)
	// 创建 FollowHandler 实例
	followHandler := handler.NewFollowHandler(followService)

	// 创建 Gin 引擎
	router := gin.Default()

	// 注册 User 相关的路由
	router.POST("/user", userHandler.CreateUserHandler)
	router.PUT("/user/:id", userHandler.UpdateUserHandler)
	router.DELETE("/user/:id", userHandler.DeleteUserHandler)
	router.GET("/user/:id", userHandler.GetUserHandler)

	router.POST("/follow", followHandler.CreateFollowHandler)
	router.PUT("/follow/:id", followHandler.UpdateFollowHandler)
	router.DELETE("/follow/:id", followHandler.DeleteFollowHandler)
	router.GET("/follow/:id", followHandler.GetFollowHandler)
	router.GET("/followers/:id", followHandler.GetFollowersByFollowingIDHandler)

	go func() {
		// 启动服务器
		err = router.Run(":8081")
		if err != nil {
			return
		}
	}()

	// grpc server start
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server listening on port 8001")

	// 启动 gRPC 服务
	s := grpc.NewServer()
	follow.RegisterFollowServiceServer(s, follow_server.NewFollowServiceServerImpl(followService))

	go func() {
		log.Println("gRPC server starting to serve...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	select {}

}
