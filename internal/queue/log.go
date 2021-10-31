package queue

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func (c *Consumer) LogCreateUser(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{BrokerAddress},
		Topic:   Topic,
		GroupID: "my-group",
	})

	for {
		msg, _ := reader.ReadMessage(ctx)
		fmt.Println("!received: ", string(msg.Value))
	}
}
