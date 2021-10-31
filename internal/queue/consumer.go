package queue

import "github.com/krupyansky/user-manager/internal/repository"

const (
	Topic         = "create-user-log"
	BrokerAddress = "localhost:9092"
)

type Consumer struct {
	repoClickHouse *repository.ClickHouse
}

func NewConsumer(repoClickHouse *repository.ClickHouse) *Consumer {
	return &Consumer{repoClickHouse: repoClickHouse}
}
