package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type stubAddr struct {
	address string
}

func (a *stubAddr) Network() string {
	return "tcp"
}

func (a *stubAddr) String() string {
	return a.address
}

type testDependencies struct {
	ctrl             *gomock.Controller
	mockPriceService *mocks.MockPriceService
	mockLogger       *mocks.StubLogger
	mockConn         *mocks.MockWebSocketConn
	stubbedAddr      *stubAddr
	handler          *LivePricesHandler
}

func setup(t *testing.T) *testDependencies {
	ctrl := gomock.NewController(t)
	mockPriceService := mocks.NewMockPriceService(ctrl)
	mockLogger := &mocks.StubLogger{}
	mockConn := mocks.NewMockWebSocketConn(ctrl)
	stubbedAddr := &stubAddr{address: "127.0.0.1:12345"}
	handler := NewLivePricesHandler(mockPriceService, mockLogger)
	return &testDependencies{
		ctrl:             ctrl,
		mockPriceService: mockPriceService,
		mockLogger:       mockLogger,
		mockConn:         mockConn,
		stubbedAddr:      stubbedAddr,
		handler:          handler,
	}
}

func TestHandleConnection_ValidSubscription(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(deps.stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Subscribe,
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	gomock.InOrder(
		deps.mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil),
		deps.mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	deps.mockPriceService.EXPECT().AddClient(deps.mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	deps.mockPriceService.EXPECT().Subscribe(deps.mockConn, subMsg.Stock).Return(nil)
	deps.mockPriceService.EXPECT().RemoveClient(deps.mockConn)
	deps.mockConn.EXPECT().Close().Return(nil)

	deps.handler.handleConnection(nil, deps.mockConn)
}

func TestHandleConnection_InvalidMessageFormat(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(deps.stubbedAddr).AnyTimes()

	invalidMessage := []byte("invalid json")

	gomock.InOrder(
		deps.mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, invalidMessage, nil),
		deps.mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	deps.mockPriceService.EXPECT().AddClient(deps.mockConn)

	expectedErrorMessage := domain.ErrorMessage{
		Type:    "error",
		Message: "Invalid message format.",
	}
	expectedErrorBytes, err := json.Marshal(expectedErrorMessage)
	assert.NoError(t, err)

	deps.mockConn.EXPECT().WriteMessage(websocket.TextMessage, expectedErrorBytes).Return(nil)
	deps.mockPriceService.EXPECT().RemoveClient(deps.mockConn)
	deps.mockConn.EXPECT().Close().Return(nil)

	deps.handler.handleConnection(nil, deps.mockConn)
}

func TestHandleConnection_UnsupportedStockSymbol(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(deps.stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Subscribe,
		Stock:  domain.Stock("UNSUPPORTED-STOCK"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	gomock.InOrder(
		deps.mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil),
		deps.mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	deps.mockPriceService.EXPECT().AddClient(deps.mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return false }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	expectedErrorMessage := domain.ErrorMessage{
		Type:    "error",
		Message: "Unsupported stock symbol",
	}
	expectedErrorBytes, err := json.Marshal(expectedErrorMessage)
	assert.NoError(t, err)

	deps.mockConn.EXPECT().WriteMessage(websocket.TextMessage, expectedErrorBytes).Return(nil)
	deps.mockPriceService.EXPECT().RemoveClient(deps.mockConn)
	deps.mockConn.EXPECT().Close().Return(nil)

	deps.handler.handleConnection(nil, deps.mockConn)
}

func TestHandleConnection_UnknownAction(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(deps.stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: "unknown_action",
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	gomock.InOrder(
		deps.mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil),
		deps.mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	deps.mockPriceService.EXPECT().AddClient(deps.mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	expectedErrorMessage := domain.ErrorMessage{
		Type:    "error",
		Message: "Unknown action",
	}
	expectedErrorBytes, err := json.Marshal(expectedErrorMessage)
	assert.NoError(t, err)

	deps.mockConn.EXPECT().WriteMessage(websocket.TextMessage, expectedErrorBytes).Return(nil)
	deps.mockPriceService.EXPECT().RemoveClient(deps.mockConn)
	deps.mockConn.EXPECT().Close().Return(nil)

	deps.handler.handleConnection(nil, deps.mockConn)
}

func TestHandleConnection_ReadMessageError(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	deps.mockConn.EXPECT().RemoteAddr().Return(deps.stubbedAddr).AnyTimes()

	deps.mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("read error"))
	deps.mockPriceService.EXPECT().AddClient(deps.mockConn)
	deps.mockPriceService.EXPECT().RemoveClient(deps.mockConn)
	deps.mockConn.EXPECT().Close().Return(nil)

	deps.handler.handleConnection(nil, deps.mockConn)
}

func TestHandleConnection_SubscribeError(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	deps.mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Subscribe,
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	deps.mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil)
	deps.mockPriceService.EXPECT().AddClient(deps.mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	subscribeErr := fmt.Errorf("subscribe error")
	deps.mockPriceService.EXPECT().Subscribe(deps.mockConn, subMsg.Stock).Return(subscribeErr)
	deps.mockPriceService.EXPECT().RemoveClient(deps.mockConn)
	deps.mockConn.EXPECT().Close().Return(nil)

	deps.handler.handleConnection(ctx, deps.mockConn)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandleConnection_UnsubscribeError(t *testing.T) {
	deps := setup(t)
	defer deps.ctrl.Finish()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	deps.mockConn.EXPECT().RemoteAddr().Return(deps.stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Unsubscribe,
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	deps.mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil)
	deps.mockPriceService.EXPECT().AddClient(deps.mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	unsubscribeErr := fmt.Errorf("unsubscribe error")
	deps.mockPriceService.EXPECT().Unsubscribe(deps.mockConn, subMsg.Stock).Return(unsubscribeErr)
	deps.mockPriceService.EXPECT().RemoveClient(deps.mockConn)
	deps.mockConn.EXPECT().Close().Return(nil)

	deps.handler.handleConnection(ctx, deps.mockConn)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
