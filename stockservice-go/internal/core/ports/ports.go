package ports

import (
	"context"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/models"
)

type Logger interface {
	Errorf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

type Consumer interface {
	Start(ctx context.Context) error
	SetListener(handlePriceEvent func(event *models.PriceEvent) error,
	)
}

type PriceEventListener interface {
	OnPriceEvent(event *models.PriceEvent) error
}

type Server interface {
	BroadcastPriceEvent(event *models.PriceEvent)
}
