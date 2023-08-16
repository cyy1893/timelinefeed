package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"relationProject/models"
	"relationProject/services"
	"strconv"
)

type FollowHandler struct {
	followService *services.FollowService
}

func NewFollowHandler(followService *services.FollowService) *FollowHandler {
	return &FollowHandler{followService: followService}
}

// CreateFollowHandler 创建 Follow 记录的处理程序
func (h *FollowHandler) CreateFollowHandler(c *gin.Context) {
	var follow models.Follow
	if err := c.ShouldBindQuery(&follow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.followService.CreateFollow(&follow); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create follow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follow created successfully"})
}

// UpdateFollowHandler 更新 Follow 记录的处理程序
func (h *FollowHandler) UpdateFollowHandler(c *gin.Context) {
	id := c.Param("id")
	followID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follow ID"})
		return
	}

	var follow models.Follow
	if err := c.ShouldBindQuery(&follow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	follow.ID = uint(followID)

	if err := h.followService.UpdateFollow(&follow); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update follow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follow updated successfully"})
}

// DeleteFollowHandler 删除 Follow 记录的处理程序
func (h *FollowHandler) DeleteFollowHandler(c *gin.Context) {
	id := c.Param("id")
	followID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follow ID"})
		return
	}

	if err := h.followService.DeleteFollowByID(uint(followID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete follow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Follow deleted successfully"})
}

// GetFollowHandler 获取 Follow 记录的处理程序
func (h *FollowHandler) GetFollowHandler(c *gin.Context) {
	id := c.Param("id")
	followID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid follow ID"})
		return
	}

	follow, err := h.followService.GetFollowByID(uint(followID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get follow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"follow": follow})
}

// GetFollowersByFollowingIDHandler 根据 FollowingID 获取 Followers 的处理程序
func (h *FollowHandler) GetFollowersByFollowingIDHandler(c *gin.Context) {
	id := c.Param("id")
	followingID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid following ID"})
		return
	}

	followers, err := h.followService.GetFollowersByFollowingID(uint(followingID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get followers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"followers": followers})
}
