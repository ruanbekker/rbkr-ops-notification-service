package kafka

import (
	"context"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer struct {
	client *kgo.Client
}

func NewConsumer(brokers string) *Consumer {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers),
		kgo.ConsumeTopics("order_reserved", "order_failed"),
		kgo.ConsumerGroup("notification-group"),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if err != nil {
		log.Fatalf("failed to create kafka consumer: %v", err)
	}

	return &Consumer{client: client}
}

func (c *Consumer) Start(handler func(topic string, msg []byte)) {
	for {
		fetches := c.client.PollFetches(context.Background())

		if errs := fetches.Errors(); len(errs) > 0 {
			for _, err := range errs {
				log.Println("kafka error:", err)
			}
			continue
		}

		fetches.EachRecord(func(record *kgo.Record) {
			log.Printf("received [%s]: %s\n", record.Topic, string(record.Value))
			handler(record.Topic, record.Value)
		})
	}
}
