package ports

import (
	"context"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain"
)

type Consumer interface {
	Start(ctx context.Context) error
	SetListener(handlePriceEvent func(event *domain.PriceEvent) error,
	)
}

type PriceEventListener interface {
	OnPriceEvent(event *domain.PriceEvent) error
}
