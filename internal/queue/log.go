package queue

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func (c *Consumer) LogCreateUser(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{BrokerAddress},
		Topic:   Topic,
		GroupID: "my-group",
	})

	for {
		msg, _ := reader.ReadMessage(ctx)
		err := c.repoClickHouse.SaveCreatedUserLog(string(msg.Value))
		if err != nil {
			log.Fatalf("save created user log err: %s", err)
		}
	}
}
