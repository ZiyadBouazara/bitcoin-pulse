package services

import (
	"context"
	"fmt"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/models"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/ports"
)

type PriceService struct {
	server   ports.Server
	consumer ports.Consumer
	logger   ports.Logger
}

func NewPriceService(server ports.Server, consumer ports.Consumer, logger ports.Logger) *PriceService {
	return &PriceService{
		server:   server,
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

func (ps *PriceService) handlePriceEvent(event *models.PriceEvent) error {
	if event == nil {
		return fmt.Errorf("received a nil PriceEvent")
	}
	ps.server.BroadcastPriceEvent(event)
	return nil
}
