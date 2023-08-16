package handlers

import (
	"feedProject/grpc/comment"
	"feedProject/grpc/follow"
	"feedProject/models"
	mysqlRepository "feedProject/repositories/mysql"
	"feedProject/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"time"
)

type FeedHandler struct {
	feedService *services.FeedService
}

func NewFeedHandler(feedRepo *mysqlRepository.Repository) (*FeedHandler, error) {
	// 建立与 gRPC 服务器的连接
	//followConn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	followConn, err := grpc.Dial("followService.default.svc.cluster.local:8001", grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	//commentConn, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	commentConn, err := grpc.Dial("commentService.default.svc.cluster.local:8002", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	feedService := services.NewFeedService(feedRepo, followConn, commentConn)
	return &FeedHandler{
		feedService: feedService,
	}, nil
}

func (h *FeedHandler) CreateFeed(c *gin.Context) {
	feed := &models.Feed{
		// 从 URL 参数中获取需要的字段值
		Content:     c.Query("content"),
		PublisherID: getIntQuery(c, "publisher_id"),
		PublishTime: getTimeQuery(c, "publish_time"),
	}

	if err := h.feedService.CreateFeed(feed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, feed)
}

func (h *FeedHandler) UpdateFeed(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	feed, err := h.feedService.GetFeedByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feed not found"})
		return
	}

	// 更新字段值
	feed.Content = c.Query("content")
	feed.PublisherID = getIntQuery(c, "publisher_id")
	feed.PublishTime = getTimeQuery(c, "publish_time")

	if err := h.feedService.UpdateFeed(feed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feed)
}

func (h *FeedHandler) DeleteFeed(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.feedService.DeleteFeedByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Feed deleted"})
}

func (h *FeedHandler) GetFeedByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	feed, err := h.feedService.GetFeedByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feed not found"})
		return
	}

	c.JSON(http.StatusOK, feed)
}

func (h *FeedHandler) GetFeedsByPusherID(c *gin.Context) {
	pusherID := getIntQuery(c, "pusher_id")
	page := getIntQuery(c, "page")
	pageSize := getIntQuery(c, "page_size")
	feeds, err := h.feedService.GetFeedsByPusherID(pusherID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feeds)
}

func (h *FeedHandler) GetFollowersByFollowingID(c *gin.Context) {
	followingID := getIntQuery(c, "following_id")
	followers, err := h.feedService.GetFollowersByFollowingID(c, &follow.FollowingIDRequest{FollowingId: uint32(followingID)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, followers)
}

// GetCommentCountByFeedID 获取指定 Feed 的评论数
func (h *FeedHandler) GetCommentCountByFeedID(c *gin.Context) {
	feedID, err := strconv.Atoi(c.Param("feed_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	count, err := h.feedService.GetCommentCountByFeedID(c, &comment.CommentCountRequest{FeedId: int32(feedID)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, count)
}

// 从 URL 参数中获取整型字段值
func getIntQuery(c *gin.Context, paramName string) int {
	param := c.Query(paramName)
	value, _ := strconv.Atoi(param)
	return value
}

// 从 URL 参数中获取时间字段值
func getTimeQuery(c *gin.Context, paramName string) time.Time {
	param := c.Query(paramName)
	// 解析时间字符串为 time.Time 类型，需要根据实际时间格式进行调整
	t, _ := time.Parse("2006-01-02 15:04:05", param)
	return t
}
