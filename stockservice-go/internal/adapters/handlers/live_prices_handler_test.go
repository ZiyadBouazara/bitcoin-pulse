package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
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

func setup(t *testing.T) (*gomock.Controller, *mocks.MockPriceService, *mocks.StubLogger, *mocks.MockWebSocketConn, *stubAddr, *LivePricesHandler) {
	ctrl := gomock.NewController(t)
	mockPriceService := mocks.NewMockPriceService(ctrl)
	mockLogger := &mocks.StubLogger{}
	mockConn := mocks.NewMockWebSocketConn(ctrl)
	stubbedAddr := &stubAddr{address: "127.0.0.1:12345"}
	handler := NewLivePricesHandler(mockPriceService, mockLogger)
	return ctrl, mockPriceService, mockLogger, mockConn, stubbedAddr, handler
}

func TestHandleConnection_ValidSubscription(t *testing.T) {
	ctrl, mockPriceService, _, mockConn, stubbedAddr, handler := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Subscribe,
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	gomock.InOrder(
		mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil),
		mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	mockPriceService.EXPECT().AddClient(mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	mockPriceService.EXPECT().Subscribe(mockConn, subMsg.Stock).Return(nil)
	mockPriceService.EXPECT().RemoveClient(mockConn)
	mockConn.EXPECT().Close().Return(nil)

	handler.handleConnection(nil, mockConn)
}

func TestHandleConnection_InvalidMessageFormat(t *testing.T) {
	ctrl, mockPriceService, _, mockConn, stubbedAddr, handler := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(stubbedAddr).AnyTimes()

	invalidMessage := []byte("invalid json")

	gomock.InOrder(
		mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, invalidMessage, nil),
		mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	mockPriceService.EXPECT().AddClient(mockConn)

	expectedErrorMessage := domain.ErrorMessage{
		Type:    "error",
		Message: "Invalid message format.",
	}
	expectedErrorBytes, err := json.Marshal(expectedErrorMessage)
	assert.NoError(t, err)

	mockConn.EXPECT().WriteMessage(websocket.TextMessage, expectedErrorBytes).Return(nil)
	mockPriceService.EXPECT().RemoveClient(mockConn)
	mockConn.EXPECT().Close().Return(nil)

	handler.handleConnection(nil, mockConn)
}

func TestHandleConnection_UnsupportedStockSymbol(t *testing.T) {
	ctrl, mockPriceService, _, mockConn, stubbedAddr, handler := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Subscribe,
		Stock:  domain.Stock("UNSUPPORTED-STOCK"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	gomock.InOrder(
		mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil),
		mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	mockPriceService.EXPECT().AddClient(mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return false }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	expectedErrorMessage := domain.ErrorMessage{
		Type:    "error",
		Message: "Unsupported stock symbol",
	}
	expectedErrorBytes, err := json.Marshal(expectedErrorMessage)
	assert.NoError(t, err)

	mockConn.EXPECT().WriteMessage(websocket.TextMessage, expectedErrorBytes).Return(nil)
	mockPriceService.EXPECT().RemoveClient(mockConn)
	mockConn.EXPECT().Close().Return(nil)

	handler.handleConnection(nil, mockConn)
}

func TestHandleConnection_UnknownAction(t *testing.T) {
	ctrl, mockPriceService, _, mockConn, stubbedAddr, handler := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: "unknown_action",
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	gomock.InOrder(
		mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil),
		mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("EOF")),
	)

	mockPriceService.EXPECT().AddClient(mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	expectedErrorMessage := domain.ErrorMessage{
		Type:    "error",
		Message: "Unknown action",
	}
	expectedErrorBytes, err := json.Marshal(expectedErrorMessage)
	assert.NoError(t, err)

	mockConn.EXPECT().WriteMessage(websocket.TextMessage, expectedErrorBytes).Return(nil)
	mockPriceService.EXPECT().RemoveClient(mockConn)
	mockConn.EXPECT().Close().Return(nil)

	handler.handleConnection(nil, mockConn)
}

func TestHandleConnection_ReadMessageError(t *testing.T) {
	ctrl, mockPriceService, _, mockConn, stubbedAddr, handler := setup(t)
	defer ctrl.Finish()

	mockConn.EXPECT().RemoteAddr().Return(stubbedAddr).AnyTimes()

	mockConn.EXPECT().ReadMessage().Return(0, nil, fmt.Errorf("read error"))
	mockPriceService.EXPECT().AddClient(mockConn)
	mockPriceService.EXPECT().RemoveClient(mockConn)
	mockConn.EXPECT().Close().Return(nil)

	handler.handleConnection(nil, mockConn)
}

func TestHandleConnection_SubscribeError(t *testing.T) {
	ctrl, mockPriceService, _, mockConn, _, handler := setup(t)
	defer ctrl.Finish()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	mockConn.EXPECT().RemoteAddr().Return(nil).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Subscribe,
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil)
	mockPriceService.EXPECT().AddClient(mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	subscribeErr := fmt.Errorf("subscribe error")
	mockPriceService.EXPECT().Subscribe(mockConn, subMsg.Stock).Return(subscribeErr)
	mockPriceService.EXPECT().RemoveClient(mockConn)
	mockConn.EXPECT().Close().Return(nil)

	handler.handleConnection(ctx, mockConn)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandleConnection_UnsubscribeError(t *testing.T) {
	ctrl, mockPriceService, _, mockConn, stubbedAddr, handler := setup(t)
	defer ctrl.Finish()

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	mockConn.EXPECT().RemoteAddr().Return(stubbedAddr).AnyTimes()

	subMsg := domain.SubscriptionMessage{
		Action: domain.Unsubscribe,
		Stock:  domain.Stock("BTC-USD"),
	}
	messageBytes, err := json.Marshal(subMsg)
	assert.NoError(t, err)

	mockConn.EXPECT().ReadMessage().Return(websocket.TextMessage, messageBytes, nil)
	mockPriceService.EXPECT().AddClient(mockConn)

	originalIsSupportedStock := domain.IsSupportedStock
	domain.IsSupportedStock = func(stock string) bool { return true }
	defer func() { domain.IsSupportedStock = originalIsSupportedStock }()

	unsubscribeErr := fmt.Errorf("unsubscribe error")
	mockPriceService.EXPECT().Unsubscribe(mockConn, subMsg.Stock).Return(unsubscribeErr)
	mockPriceService.EXPECT().RemoveClient(mockConn)
	mockConn.EXPECT().Close().Return(nil)

	handler.handleConnection(ctx, mockConn)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
