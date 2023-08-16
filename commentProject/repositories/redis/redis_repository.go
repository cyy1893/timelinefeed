package redisRepository

//
import (
	"commentProject/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	// RedisKeyCommentPrefix Redis 中 Comment 的 Key 前缀
	RedisKeyCommentPrefix      = "comment:%d"
	RedisKeyCommentCountPrefix = "comment_count:%d"
)

type Repository struct {
	redisClient *redis.Client
}

func NewRepository(redisClient *redis.Client) *Repository {
	return &Repository{
		redisClient: redisClient,
	}
}

// CreateComment 创建 Comment 记录并存储到 Redis
func (r *Repository) CreateComment(ctx context.Context, comment *models.Comment, ttl time.Duration) error {
	// 将 Comment 序列化为 JSON
	data, err := json.Marshal(comment)
	if err != nil {
		return err
	}

	// 存储 Comment 到 Redis
	key := fmt.Sprintf(RedisKeyCommentPrefix, comment.ID)
	err = r.redisClient.Set(ctx, key, data, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

// UpdateComment 更新 Comment 记录并更新 Redis
func (r *Repository) UpdateComment(ctx context.Context, comment *models.Comment, ttl time.Duration) error {
	// 将 Comment 序列化为 JSON
	data, err := json.Marshal(comment)
	if err != nil {
		return err
	}

	// 更新 Redis 中的 Comment
	key := fmt.Sprintf(RedisKeyCommentPrefix, comment.ID)
	err = r.redisClient.Set(ctx, key, data, ttl).Err()
	if err != nil {
		return err
	}

	return nil
}

// DeleteComment 删除 Comment 记录并删除 Redis 中的数据
func (r *Repository) DeleteComment(ctx context.Context, commentID int) error {
	// 删除 Redis 中的 Comment
	key := fmt.Sprintf(RedisKeyCommentPrefix, commentID)
	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetCommentByID 根据 CommentID 从 Redis 获取 Comment 记录
func (r *Repository) GetCommentByID(ctx context.Context, commentID int) (*models.Comment, error) {
	// 从 Redis 获取 Comment 数据
	key := fmt.Sprintf(RedisKeyCommentPrefix, commentID)
	data, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	// 反序列化 JSON 数据为 Comment 对象
	var comment models.Comment
	err = json.Unmarshal(data, &comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// 创建CommentCount记录
func (r *Repository) CreateCommentCount(ctx context.Context, commentCount *models.CommentCount, ttl time.Duration) error {
	key := fmt.Sprintf(RedisKeyCommentCountPrefix, commentCount.FeedID)
	data, err := json.Marshal(commentCount)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, data, ttl).Err()
}

// 获取CommentCount记录
func (r *Repository) GetCommentCount(ctx context.Context, id int) (*models.CommentCount, error) {
	key := fmt.Sprintf(RedisKeyCommentCountPrefix, id)
	data, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return &models.CommentCount{}, err
	}
	var commentCount models.CommentCount
	err = json.Unmarshal(data, &commentCount)
	if err != nil {
		return &models.CommentCount{}, err
	}
	return &commentCount, nil
}

// 更新CommentCount记录
func (r *Repository) UpdateCommentCount(ctx context.Context, commentCount *models.CommentCount, ttl time.Duration) error {
	key := fmt.Sprintf(RedisKeyCommentCountPrefix, commentCount.FeedID)
	data, err := json.Marshal(commentCount)
	if err != nil {
		return err
	}
	return r.redisClient.Set(ctx, key, data, ttl).Err()
}

// 删除CommentCount记录
func (r *Repository) DeleteCommentCount(ctx context.Context, id int) error {
	key := fmt.Sprintf(RedisKeyCommentCountPrefix, id)
	return r.redisClient.Del(ctx, key).Err()
}
