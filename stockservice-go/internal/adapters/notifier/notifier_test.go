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

type testDependencies struct {
	ctrl       *gomock.Controller
	stubLogger *mocks.StubLogger
	mockConn   *mocks.MockWebSocketConn
	notifier   *Notifier
}

func setup(t *testing.T) *testDependencies {
	ctrl := gomock.NewController(t)
	stubLogger := &mocks.StubLogger{}
	mockConn := mocks.NewMockWebSocketConn(ctrl)
	notifier := NewNotifier(stubLogger)
	return &testDependencies{
		ctrl:       ctrl,
		stubLogger: stubLogger,
		mockConn:   mockConn,
		notifier:   notifier,
	}
}

func TestNotifier_AddClient(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.notifier.AddClient(deps.mockConn)

	assert.Contains(t, deps.notifier.GetConnections(), deps.mockConn)
}

func TestNotifier_RemoveClient(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.notifier.AddClient(deps.mockConn)
	deps.notifier.RemoveClient(deps.mockConn)

	assert.NotContains(t, deps.notifier.GetConnections(), deps.mockConn)
}

func TestNotifier_Subscribe(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	err := deps.notifier.Subscribe(deps.mockConn, aStock)

	assert.NoError(t, err)
	assert.Contains(t, deps.notifier.GetSubscriptions(aStock), deps.mockConn)
}

func TestNotifier_Unsubscribe(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	_ = deps.notifier.Subscribe(deps.mockConn, aStock)

	err := deps.notifier.Unsubscribe(deps.mockConn, aStock)

	assert.NoError(t, err)
	assert.NotContains(t, deps.notifier.GetSubscriptions(aStock), deps.mockConn)
}

func TestNotifier_Broadcast(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	_ = deps.notifier.Subscribe(deps.mockConn, aStock)

	event := &domain.PriceEvent{
		ProductID: aStock,
		Price:     50000.00,
	}

	msg, err := json.Marshal(event)
	assert.NoError(t, err)
	deps.mockConn.EXPECT().WriteMessage(websocket.TextMessage, msg).Return(nil).Times(1)

	err = deps.notifier.Broadcast(event)

	assert.NoError(t, err)
}

func TestNotifier_Broadcast_WriteMessageError(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()
	deps.mockConn.EXPECT().Close().Return(nil)
	_ = deps.notifier.Subscribe(deps.mockConn, aStock)

	event := &domain.PriceEvent{
		ProductID: aStock,
		Price:     50000.00,
	}

	msg, err := json.Marshal(event)
	assert.NoError(t, err)
	writeErr := fmt.Errorf("write error")
	deps.mockConn.EXPECT().WriteMessage(websocket.TextMessage, msg).Return(writeErr).Times(1)

	err = deps.notifier.Broadcast(event)

	assert.NoError(t, err)
	assert.NotContains(t, deps.notifier.GetConnections(), deps.mockConn)
}
