package services

import (
	"context"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/ports"
)

type PriceService struct {
	notifier ports.Notifier
	consumer ports.Consumer
	logger   ports.Logger
}

func NewPriceService(notifier ports.Notifier, consumer ports.Consumer, logger ports.Logger) *PriceService {
	return &PriceService{
		notifier: notifier,
		consumer: consumer,
		logger:   logger,
	}
}

func (ps *PriceService) StartConsuming(ctx context.Context) {
	ps.consumer.SetListener(ps.notifier.Broadcast)

	if err := ps.consumer.Start(ctx); err != nil {
		ps.logger.Errorf("BitcoinPriceConsumer exited with error: %v", err)
	} else {
		ps.logger.Info("BitcoinPriceConsumer exited")
	}
}

func (ps *PriceService) AddClient(ws ports.WebSocketConn) {
	ps.notifier.AddClient(ws)
}

func (ps *PriceService) RemoveClient(ws ports.WebSocketConn) {
	ps.notifier.RemoveClient(ws)
}

func (ps *PriceService) Subscribe(ws ports.WebSocketConn, stock domain.Stock) error {
	err := ps.notifier.Subscribe(ws, stock)
	if err != nil {
		ps.logger.Errorf("error subscribing Client: %v", ws.RemoteAddr())
	}
	return nil
}

func (ps *PriceService) Unsubscribe(ws ports.WebSocketConn, stock domain.Stock) error {
	err := ps.notifier.Unsubscribe(ws, stock)
	if err != nil {
		ps.logger.Errorf("error unsubscribing Client: %v", ws.RemoteAddr())
	}
	return nil
}
