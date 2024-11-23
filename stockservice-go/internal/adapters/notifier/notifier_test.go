package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/mocks"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

var aStock = domain.Stock("BTC-USD")

func setup(t *testing.T) (*gomock.Controller, *mocks.StubLogger, *mocks.MockWebSocketConn, *Notifier) {
	ctrl := gomock.NewController(t)
	stubLogger := &mocks.StubLogger{}
	mockConn := mocks.NewMockWebSocketConn(ctrl)
	notifier := NewNotifier(stubLogger)
	return ctrl, stubLogger, mockConn, notifier
}

func TestNotifier_AddClient(t *testing.T) {
	ctrl, _, mockConn, notifier := setup(t)
	defer ctrl.Finish()

	notifier.AddClient(mockConn)

	assert.Contains(t, notifier.GetConnections(), mockConn)
}

func TestNotifier_RemoveClient(t *testing.T) {
	ctrl, _, mockConn, notifier := setup(t)
	defer ctrl.Finish()

	notifier.AddClient(mockConn)
	notifier.RemoveClient(mockConn)

	assert.NotContains(t, notifier.GetConnections(), mockConn)
}

func TestNotifier_Subscribe(t *testing.T) {
	ctrl, _, mockConn, notifier := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	err := notifier.Subscribe(mockConn, aStock)

	assert.NoError(t, err)
	assert.Contains(t, notifier.GetSubscriptions(aStock), mockConn)
}

func TestNotifier_Unsubscribe(t *testing.T) {
	ctrl, _, mockConn, notifier := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	_ = notifier.Subscribe(mockConn, aStock)

	err := notifier.Unsubscribe(mockConn, aStock)

	assert.NoError(t, err)
	assert.NotContains(t, notifier.GetSubscriptions(aStock), mockConn)
}

func TestNotifier_Broadcast(t *testing.T) {
	ctrl, _, mockConn, notifier := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	_ = notifier.Subscribe(mockConn, aStock)

	event := &domain.PriceEvent{
		ProductID: aStock,
		Price:     50000.00,
	}

	msg, err := json.Marshal(event)
	assert.NoError(t, err)
	mockConn.EXPECT().WriteMessage(websocket.TextMessage, msg).Return(nil).Times(1)

	err = notifier.Broadcast(event)

	assert.NoError(t, err)
}

func TestNotifier_Broadcast_WriteMessageError(t *testing.T) {
	ctrl, _, mockConn, notifier := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	mockConn.EXPECT().Close().Return(nil)
	_ = notifier.Subscribe(mockConn, aStock)

	event := &domain.PriceEvent{
		ProductID: aStock,
		Price:     50000.00,
	}

	msg, err := json.Marshal(event)
	assert.NoError(t, err)
	writeErr := fmt.Errorf("write error")
	mockConn.EXPECT().WriteMessage(websocket.TextMessage, msg).Return(writeErr).Times(1)

	err = notifier.Broadcast(event)

	assert.NoError(t, err)
	assert.NotContains(t, notifier.GetConnections(), mockConn)
}
