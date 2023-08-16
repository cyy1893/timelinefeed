package main

import (
	"context"
	"feedProject/config"
	"feedProject/handlers"
	mysqlRepository "feedProject/repositories/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	// 初始化 Gin 引擎
	r := gin.Default()

	// 创建 FeedRepository 实例
	feedRepo := mysqlRepository.NewFeedRepository(db)

	// 创建 FeedHandler 实例
	handler, err := handlers.NewFeedHandler(feedRepo)
	if err != nil {
		panic(err)
	}

	// 设置路由和处理程序
	r.POST("/feeds", handler.CreateFeed)
	r.PUT("/feeds/:id", handler.UpdateFeed)
	r.DELETE("/feeds/:id", handler.DeleteFeed)
	r.GET("/feeds/:id", handler.GetFeedByID)
	r.GET("/followers/:id", handler.GetFollowersByFollowingID)
	r.GET("/commentCount/:feed_id", handler.GetCommentCountByFeedID)
	r.GET("/feeds", handler.GetFeedsByPusherID)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败：%s\n", err.Error())
		}
	}()

	log.Println("服务器已启动")

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	log.Println("接收到中断信号...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("服务器关闭失败：%s\n", err.Error())
	}

	log.Println("服务器已停止")
}
