package queue

const (
	Topic         = "create-user-log"
	BrokerAddress = "localhost:9092"
)

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}
