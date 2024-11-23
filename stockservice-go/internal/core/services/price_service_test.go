package services

import (
	"context"
	"errors"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func setup(t *testing.T) (*gomock.Controller, *mocks.MockNotifier, *mocks.MockWebSocketConn, *mocks.MockConsumer, *PriceService) {
	ctrl := gomock.NewController(t)
	mockNotifier := mocks.NewMockNotifier(ctrl)
	mockConn := mocks.NewMockWebSocketConn(ctrl)
	mockConsumer := mocks.NewMockConsumer(ctrl)
	stubLogger := &mocks.StubLogger{}
	service := NewPriceService(mockNotifier, mockConsumer, stubLogger)

	return ctrl, mockNotifier, mockConn, mockConsumer, service
}

func TestPriceService_StartConsuming_WithError(t *testing.T) {
	ctrl, _, _, mockConsumer, priceService := setup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	startErr := errors.New("consumer failed to start")

	mockConsumer.EXPECT().SetListener(gomock.Any())
	mockConsumer.EXPECT().Start(ctx).Return(startErr)

	priceService.StartConsuming(ctx)
}

func TestPriceService_StartConsuming_WithNoError(t *testing.T) {
	ctrl, _, _, mockConsumer, priceService := setup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockConsumer.EXPECT().SetListener(gomock.Any())
	mockConsumer.EXPECT().Start(ctx).Return(nil)

	priceService.StartConsuming(ctx)
}

func TestPriceService_AddClient(t *testing.T) {
	ctrl, mockNotifier, mockConn, _, priceService := setup(t)
	defer ctrl.Finish()

	mockNotifier.EXPECT().AddClient(mockConn)

	priceService.AddClient(mockConn)
}

func TestPriceService_RemoveClient(t *testing.T) {
	ctrl, mockNotifier, mockConn, _, priceService := setup(t)
	defer ctrl.Finish()

	mockNotifier.EXPECT().RemoveClient(mockConn)

	priceService.RemoveClient(mockConn)
}

func TestPriceService_Subscribe(t *testing.T) {
	ctrl, mockNotifier, mockConn, _, priceService := setup(t)
	defer ctrl.Finish()

	stock := domain.Stock("BTC-USD")
	mockNotifier.EXPECT().Subscribe(mockConn, stock).Return(nil)

	err := priceService.Subscribe(mockConn, stock)
	assert.NoError(t, err)
}

func TestPriceService_Unsubscribe(t *testing.T) {
	ctrl, mockNotifier, mockConn, _, priceService := setup(t)
	defer ctrl.Finish()

	stock := domain.Stock("BTC-USD")
	mockNotifier.EXPECT().Unsubscribe(mockConn, stock).Return(nil)

	err := priceService.Unsubscribe(mockConn, stock)
	assert.NoError(t, err)
}
