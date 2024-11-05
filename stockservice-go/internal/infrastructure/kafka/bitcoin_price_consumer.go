package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/infrastructure/dtos"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type BitcoinPriceConsumer struct {
	reader  *kafka.Reader
	handler func(event *domain.PriceEvent) error
	logger  *logrus.Logger
}

func NewBitcoinPriceConsumer(brokerURL, topic, groupID string, logger *logrus.Logger) *BitcoinPriceConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		GroupID: groupID,
		Topic:   topic,
	})

	return &BitcoinPriceConsumer{
		reader: reader,
		logger: logger,
	}
}

func (c *BitcoinPriceConsumer) SetListener(handlePriceEvent func(event *domain.PriceEvent) error) {
	c.handler = handlePriceEvent
}

func (c *BitcoinPriceConsumer) Start(ctx context.Context) error {
	defer func() {
		if err := c.reader.Close(); err != nil {
			c.logger.Errorf("Error closing Kafka reader: %v", err)
		}
	}()
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				c.logger.Info("BitcoinPriceConsumer context canceled")
				return nil
			}
			c.logger.Errorf("Error reading message: %v", err)
			continue
		}

		c.logger.Debugf("Message received at offset %d: %s", msg.Offset, string(msg.Value))

		var eventDTO dtos.PriceEventDTO
		if err := json.Unmarshal(msg.Value, &eventDTO); err != nil {
			c.logger.Errorf("Error unmarshalling message: %v", err)
			continue
		}

		c.logger.Infof("🚀 🚀 🚀 BTC Price Event received 🚀 🚀 🚀 %s", eventDTO.FormatLog())

		priceEvent, err := dtos.ToPriceEvent(&eventDTO)
		if err != nil {
			c.logger.Errorf("Error converting PriceEventDTO -> PriceEvent message: %v", err)
			continue
		}

		if err := c.handler(priceEvent); err != nil {
			c.logger.Errorf("Error handling message: %v", err)
			continue
		}

		c.logger.Debugf("Processed message at offset %d", msg.Offset)
	}
}
