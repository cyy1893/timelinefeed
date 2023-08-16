package services

import (
	"commentProject/models"
	kafkaRepository "commentProject/repositories/kafka"
	"commentProject/repositories/mysql"
	redisRepository "commentProject/repositories/redis"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type CommentService struct {
	mysqlRepository *mysqlRepository.Repository
	redisRepository *redisRepository.Repository
	kafkaRepository *kafkaRepository.Repository
}

func NewCommentService(commentRepository *mysqlRepository.Repository, redisRepository *redisRepository.Repository, kafkaRepository *kafkaRepository.Repository) *CommentService {
	return &CommentService{
		mysqlRepository: commentRepository,
		redisRepository: redisRepository,
		kafkaRepository: kafkaRepository,
	}
}

// CreateComment 创建 Comment 记录
func (s *CommentService) CreateComment(ctx context.Context, comment *models.Comment) error {
	commentCount, err := s.redisRepository.GetCommentCount(ctx, comment.FeedID)
	if err == nil {
		commentCount.CommentCount++
		if err := s.redisRepository.CreateCommentCount(ctx, commentCount, 24*time.Hour); err != nil {
			return err
		}
	}

	err = s.kafkaRepository.SendCommentToKafka(ctx, s.kafkaRepository.Producer, comment)
	return err
}

func (s *CommentService) ConsumeCommentMessage(ctx context.Context, comment *models.Comment) {
	messageChan := make(chan string, 1000)
	go s.kafkaRepository.ReceiveCommentsFromKafka(ctx, s.kafkaRepository.Consumer, messageChan)
	go func() {
		for {
			select {
			case message := <-messageChan:
				err := s.mysqlRepository.CreateComment(comment)
				if err != nil {
					log.Fatal(message)
				}
			default:
				// ......
			}
		}
	}()
}

// UpdateComment 更新 Comment 记录
func (s *CommentService) UpdateComment(ctx context.Context, comment *models.Comment) error {
	if err := s.redisRepository.DeleteComment(ctx, int(comment.ID)); err != nil {
		return err
	}
	return s.mysqlRepository.UpdateComment(comment)
}

// DeleteComment 删除 Comment 记录
func (s *CommentService) DeleteComment(ctx context.Context, comment *models.Comment) error {
	commentCount, err := s.redisRepository.GetCommentCount(ctx, comment.FeedID)
	if err == nil {
		commentCount.CommentCount--
		if err := s.redisRepository.CreateCommentCount(ctx, commentCount, 24*time.Hour); err != nil {
			return err
		}
	}

	if err := s.redisRepository.DeleteComment(ctx, int(comment.ID)); err != nil {
		return err
	}
	return s.mysqlRepository.DeleteComment(comment)
}

// GetCommentByID 根据 CommentID 获取 Comment 记录
func (s *CommentService) GetCommentByID(ctx context.Context, commentID int) (*models.Comment, error) {
	// 优先从redis中获取记录，在获取不到的情况下从mysql中获取，并且把记录存储到redis中
	comment, err := s.redisRepository.GetCommentByID(ctx, commentID)
	if err == redis.Nil {
		comment, err = s.mysqlRepository.GetCommentByID(commentID)
		if err != nil {
			return nil, err
		}
		if comment != nil {
			if err := s.redisRepository.CreateComment(ctx, comment, 24*time.Hour); err != nil {
				return nil, err
			}
		}
		return comment, nil
	} else if err != nil {
		return nil, err
	} else {
		return comment, nil
	}
}

// GetCommentCountByFeedID 根据 FeedID 获取 CommentCount 记录的数量
func (s *CommentService) GetCommentCountByFeedID(ctx context.Context, feedID int) (int, error) {
	commentCount, err := s.redisRepository.GetCommentCount(ctx, feedID)
	if err == redis.Nil {
		count, err := s.mysqlRepository.GetCommentCountByFeedID(feedID)
		if err != nil {
			return 0, err
		}
		if count != 0 {
			if err := s.redisRepository.CreateCommentCount(ctx, commentCount, 24*time.Hour); err != nil {
				return 0, err
			}
		}
		return count, nil
	} else if err != nil {
		return 0, err
	} else {
		return commentCount.CommentCount, nil
	}
}

// GetCommentsByUserID 根据UserID 获取 Comments
func (s *CommentService) GetCommentsByUserID(ctx context.Context, userID int) ([]models.Comment, error) {
	comments, err := s.mysqlRepository.GetCommentsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
