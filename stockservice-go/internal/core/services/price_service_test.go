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

func setup(t *testing.T) (*gomock.Controller, *mocks.MockServer, *mocks.MockConsumer, *PriceService) {
	ctrl := gomock.NewController(t)
	mockServer := mocks.NewMockServer(ctrl)
	mockConsumer := mocks.NewMockConsumer(ctrl)
	stubLogger := &mocks.StubLogger{}
	service := NewPriceService(mockServer, mockConsumer, stubLogger)

	return ctrl, mockServer, mockConsumer, service
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

func TestPriceService_handlePriceEvent_WithNilEvent(t *testing.T) {
	ctrl, _, _, priceService := setup(t)
	defer ctrl.Finish()

	err := priceService.handlePriceEvent(nil)

	assert.Error(t, err)
	assert.Equal(t, "received a nil PriceEvent", err.Error())
}

func TestPriceService_handlePriceEvent_WithValidEvent(t *testing.T) {
	ctrl, mockService, _, priceService := setup(t)
	defer ctrl.Finish()

	priceEvent := &domain.PriceEvent{
		Price: 100000.0,
	}
	mockService.EXPECT().BroadcastPriceEvent(priceEvent)

	err := priceService.handlePriceEvent(priceEvent)

	assert.NoError(t, err)
}
