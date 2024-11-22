package services

import (
	"context"
	"errors"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/mocks"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func setup(t *testing.T) (*gomock.Controller, *mocks.MockNotifier, *mocks.MockConsumer, *PriceService) {
	ctrl := gomock.NewController(t)
	mockNotifier := mocks.NewMockNotifier(ctrl)
	mockConsumer := mocks.NewMockConsumer(ctrl)
	stubLogger := &mocks.StubLogger{}
	service := NewPriceService(mockNotifier, mockConsumer, stubLogger)

	return ctrl, mockNotifier, mockConsumer, service
}

func TestPriceService_StartConsuming_WithError(t *testing.T) {
	ctrl, _, mockConsumer, priceService := setup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	startErr := errors.New("consumer failed to start")

	mockConsumer.EXPECT().SetListener(gomock.Any())
	mockConsumer.EXPECT().Start(ctx).Return(startErr)

	priceService.StartConsuming(ctx)
}

func TestPriceService_StartConsuming_WithNoError(t *testing.T) {
	ctrl, _, mockConsumer, priceService := setup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockConsumer.EXPECT().SetListener(gomock.Any())
	mockConsumer.EXPECT().Start(ctx).Return(nil)

	priceService.StartConsuming(ctx)
}

func TestPriceService_AddClient(t *testing.T) {
	ctrl, mockNotifier, _, priceService := setup(t)
	defer ctrl.Finish()

	ws := &websocket.Conn{}
	mockNotifier.EXPECT().AddClient(ws)

	priceService.AddClient(ws)
}

func TestPriceService_RemoveClient(t *testing.T) {
	ctrl, mockNotifier, _, priceService := setup(t)
	defer ctrl.Finish()

	ws := &websocket.Conn{}
	mockNotifier.EXPECT().RemoveClient(ws)

	priceService.RemoveClient(ws)
}

func TestPriceService_Subscribe(t *testing.T) {
	ctrl, mockNotifier, _, priceService := setup(t)
	defer ctrl.Finish()

	ws := &websocket.Conn{}
	stock := domain.Stock("BTC-USD")
	mockNotifier.EXPECT().Subscribe(ws, stock).Return(nil)

	err := priceService.Subscribe(ws, stock)
	assert.NoError(t, err)
}

func TestPriceService_Unsubscribe(t *testing.T) {
	ctrl, mockNotifier, _, priceService := setup(t)
	defer ctrl.Finish()

	ws := &websocket.Conn{}
	stock := domain.Stock("BTC-USD")
	mockNotifier.EXPECT().Unsubscribe(ws, stock).Return(nil)

	err := priceService.Unsubscribe(ws, stock)
	assert.NoError(t, err)
}
