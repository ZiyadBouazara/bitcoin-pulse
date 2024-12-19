package kafka

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/adapters/dtos"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/mocks"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/testutils"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setupConsumer(handler func(event *domain.PriceEvent) error) *BitcoinPriceConsumer {
	return &BitcoinPriceConsumer{
		logger:  &mocks.StubLogger{},
		handler: handler,
	}
}

func createKafkaMessage(value interface{}) kafka.Message {
	messageBytes, _ := json.Marshal(value)
	return kafka.Message{
		Offset: 0,
		Value:  messageBytes,
	}
}

func TestBitcoinPriceConsumer_ProcessMessage_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handlerCalled := false
	handler := func(event *domain.PriceEvent) error {
		handlerCalled = true
		return nil
	}

	consumer := setupConsumer(handler)
	eventDTO := testutils.CreateValidPriceEventDTO()
	msg := createKafkaMessage(eventDTO)

	err := consumer.ProcessMessage(msg)

	assert.NoError(t, err)
	assert.True(t, handlerCalled, "Handler should have been called")
}

func TestBitcoinPriceConsumer_ProcessMessage_UnmarshalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	consumer := setupConsumer(func(event *domain.PriceEvent) error { return nil })
	invalidMessage := kafka.Message{
		Offset: 0,
		Value:  []byte(`invalid json`),
	}

	err := consumer.ProcessMessage(invalidMessage)

	assert.Error(t, err)
}

func TestBitcoinPriceConsumer_ProcessMessage_ConversionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	consumer := setupConsumer(func(event *domain.PriceEvent) error { return nil })
	eventDTO := dtos.PriceEventDTO{Price: ""}
	msg := createKafkaMessage(eventDTO)

	err := consumer.ProcessMessage(msg)

	assert.Error(t, err)
}

func TestBitcoinPriceConsumer_ProcessMessage_HandlerError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handlerErr := errors.New("handler failed")
	consumer := setupConsumer(func(event *domain.PriceEvent) error { return handlerErr })
	eventDTO := testutils.CreateValidPriceEventDTO()
	msg := createKafkaMessage(eventDTO)

	err := consumer.ProcessMessage(msg)

	assert.Error(t, err)
	assert.Equal(t, handlerErr, err)
}
