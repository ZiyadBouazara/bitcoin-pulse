package kafka

import (
	"context"
	"encoding/json"
	"github.com/ZiyadBouazara/stockservice-go/internal/domain"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	reader  *kafka.Reader
	handler func(event *domain.PriceEvent) error
	logger  *logrus.Logger
}

func NewConsumer(brokerURL, topic, groupID string, handler func(event *domain.PriceEvent) error, logger *logrus.Logger) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		GroupID: groupID,
		Topic:   topic,
	})

	return &Consumer{
		reader:  reader,
		handler: handler,
		logger:  logger,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	defer c.reader.Close()
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if err == context.Canceled {
				c.logger.Info("Consumer context canceled")
				return nil
			}
			c.logger.Errorf("Error reading message: %v", err)
			continue
		}

		c.logger.Infof("Message received at offset %d: %s", msg.Offset, string(msg.Value))

		var event domain.PriceEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			c.logger.Errorf("Error unmarshalling message: %v", err)
			continue
		}

		if err := c.handler(&event); err != nil {
			c.logger.Errorf("Error handling message: %v", err)
			continue
		}

		c.logger.Infof("Processed message at offset %d", msg.Offset)
	}
}
