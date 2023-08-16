package kafkaRepository

import (
	"commentProject/models"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

const (
	KafkaKeyCommentPrefix = "comment-%d"
)

type Repository struct {
	Producer *kafka.Writer
	Consumer *kafka.Reader
}

func NewKafkaRepository(brokerAddress string, topic string) *Repository {
	return &Repository{
		Producer: createKafkaProducer(brokerAddress, topic),
		Consumer: createKafkaConsumer(brokerAddress, topic),
	}
}

// 创建Kafka生产者
func createKafkaProducer(brokerAddress string, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.Hash{},
		BatchTimeout: 10 * time.Millisecond,
		BatchSize:    100,
		BatchBytes:   1048576,
	}
}

// 创建Kafka消费者
func createKafkaConsumer(brokerAddress string, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})
}

// 发送评论到Kafka
func (k *Repository) SendCommentToKafka(ctx context.Context, producer *kafka.Writer, comment *models.Comment) error {
	err := producer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(fmt.Sprintf(KafkaKeyCommentPrefix, comment.UserID)),
			Value: []byte(comment.CommentContent),
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("Comment sent to Kafka")
	return nil
}

func (k *Repository) ReceiveCommentsFromKafka(ctx context.Context, consumer *kafka.Reader, messageChan chan<- string) {
	for {
		message, err := consumer.ReadMessage(ctx)
		if err != nil {
			log.Fatal("Failed to read comment message from Kafka:", err)
		}
		messageChan <- string(message.Value)
	}
}
