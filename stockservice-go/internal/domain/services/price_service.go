package services

import (
	"context"
	"fmt"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain/ports"
	"github.com/sirupsen/logrus"
)

type PriceService struct {
	consumer ports.Consumer
	logger   *logrus.Logger
}

func NewPriceService(consumer ports.Consumer, logger *logrus.Logger) *PriceService {
	return &PriceService{
		consumer: consumer,
		logger:   logger,
	}
}

func (ps *PriceService) StartConsuming(ctx context.Context) {
	ps.consumer.SetListener(ps.handlePriceEvent)

	if err := ps.consumer.Start(ctx); err != nil {
		ps.logger.Errorf("BitcoinPriceConsumer exited with error: %v", err)
	} else {
		ps.logger.Info("BitcoinPriceConsumer exited")
	}
}

func (ps *PriceService) handlePriceEvent(event *domain.PriceEvent) error {
	if event == nil {
		return fmt.Errorf("received a nil PriceEvent")
	}
	return nil
}
