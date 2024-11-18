package services

import (
	"context"
	"errors"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func setup(t *testing.T) (*gomock.Controller, *mocks.MockConsumer, *mocks.MockLogger, *PriceService) {
	ctrl := gomock.NewController(t)
	mockConsumer := mocks.NewMockConsumer(ctrl)
	mockLogger := mocks.NewMockLogger(ctrl)
	service := NewPriceService(mockConsumer, mockLogger)

	return ctrl, mockConsumer, mockLogger, service
}

func TestPriceService_StartConsuming_WithError(t *testing.T) {
	ctrl, mockConsumer, mockLogger, priceService := setup(t)
	defer ctrl.Finish()

	ctx := context.Background()
	startErr := errors.New("consumer failed to start")

	mockConsumer.EXPECT().SetListener(gomock.Any())
	mockConsumer.EXPECT().Start(ctx).Return(startErr)
	mockLogger.EXPECT().Errorf("BitcoinPriceConsumer exited with error: %v", startErr)

	priceService.StartConsuming(ctx)
}

func TestPriceService_StartConsuming_WithNoError(t *testing.T) {
	ctrl, mockConsumer, mockLogger, priceService := setup(t)
	defer ctrl.Finish()

	ctx := context.Background()

	mockConsumer.EXPECT().SetListener(gomock.Any())
	mockConsumer.EXPECT().Start(ctx).Return(nil)
	mockLogger.EXPECT().Info("BitcoinPriceConsumer exited")

	priceService.StartConsuming(ctx)
}

func TestPriceService_handlePriceEvent_WithNilEvent(t *testing.T) {
	ctrl, _, _, priceService := setup(t)
	defer ctrl.Finish()

	err := priceService.handlePriceEvent(nil)

	assert.Error(t, err)
	assert.Equal(t, "received a nil PriceEvent", err.Error())
}

func TestPriceService_handlePriceEvent_WithValidEvent(t *testing.T) {
	ctrl, _, _, priceService := setup(t)
	defer ctrl.Finish()

	priceEvent := &domain.PriceEvent{
		Price: 100000.0,
	}

	err := priceService.handlePriceEvent(priceEvent)

	assert.NoError(t, err)
}
