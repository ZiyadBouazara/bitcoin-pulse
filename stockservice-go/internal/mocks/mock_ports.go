// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/core/ports/ports.go
//
// Generated by this command:
//
//	mockgen -source=./internal/core/ports/ports.go -destination=./internal/mocks/mock_ports.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	websocket "github.com/gorilla/websocket"
	gomock "go.uber.org/mock/gomock"
)

// MockPriceService is a mock of PriceService interface.
type MockPriceService struct {
	ctrl     *gomock.Controller
	recorder *MockPriceServiceMockRecorder
	isgomock struct{}
}

// MockPriceServiceMockRecorder is the mock recorder for MockPriceService.
type MockPriceServiceMockRecorder struct {
	mock *MockPriceService
}

// NewMockPriceService creates a new mock instance.
func NewMockPriceService(ctrl *gomock.Controller) *MockPriceService {
	mock := &MockPriceService{ctrl: ctrl}
	mock.recorder = &MockPriceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceService) EXPECT() *MockPriceServiceMockRecorder {
	return m.recorder
}

// AddClient mocks base method.
func (m *MockPriceService) AddClient(ws *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddClient", ws)
}

// AddClient indicates an expected call of AddClient.
func (mr *MockPriceServiceMockRecorder) AddClient(ws any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockPriceService)(nil).AddClient), ws)
}

// RemoveClient mocks base method.
func (m *MockPriceService) RemoveClient(ws *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveClient", ws)
}

// RemoveClient indicates an expected call of RemoveClient.
func (mr *MockPriceServiceMockRecorder) RemoveClient(ws any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveClient", reflect.TypeOf((*MockPriceService)(nil).RemoveClient), ws)
}

// StartConsuming mocks base method.
func (m *MockPriceService) StartConsuming(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartConsuming", ctx)
}

// StartConsuming indicates an expected call of StartConsuming.
func (mr *MockPriceServiceMockRecorder) StartConsuming(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartConsuming", reflect.TypeOf((*MockPriceService)(nil).StartConsuming), ctx)
}

// Subscribe mocks base method.
func (m *MockPriceService) Subscribe(ws *websocket.Conn, stock domain.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ws, stock)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockPriceServiceMockRecorder) Subscribe(ws, stock any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockPriceService)(nil).Subscribe), ws, stock)
}

// Unsubscribe mocks base method.
func (m *MockPriceService) Unsubscribe(ws *websocket.Conn, stock domain.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ws, stock)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockPriceServiceMockRecorder) Unsubscribe(ws, stock any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockPriceService)(nil).Unsubscribe), ws, stock)
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
	isgomock struct{}
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Info mocks base method.
func (m *MockLogger) Info(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), args...)
}

// Infof mocks base method.
func (m *MockLogger) Infof(format string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(format any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// MockConsumer is a mock of Consumer interface.
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
	isgomock struct{}
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer.
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance.
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// SetListener mocks base method.
func (m *MockConsumer) SetListener(handlePriceEvent func(*domain.PriceEvent) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetListener", handlePriceEvent)
}

// SetListener indicates an expected call of SetListener.
func (mr *MockConsumerMockRecorder) SetListener(handlePriceEvent any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetListener", reflect.TypeOf((*MockConsumer)(nil).SetListener), handlePriceEvent)
}

// Start mocks base method.
func (m *MockConsumer) Start(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockConsumerMockRecorder) Start(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockConsumer)(nil).Start), ctx)
}

// MockPriceEventListener is a mock of PriceEventListener interface.
type MockPriceEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockPriceEventListenerMockRecorder
	isgomock struct{}
}

// MockPriceEventListenerMockRecorder is the mock recorder for MockPriceEventListener.
type MockPriceEventListenerMockRecorder struct {
	mock *MockPriceEventListener
}

// NewMockPriceEventListener creates a new mock instance.
func NewMockPriceEventListener(ctrl *gomock.Controller) *MockPriceEventListener {
	mock := &MockPriceEventListener{ctrl: ctrl}
	mock.recorder = &MockPriceEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceEventListener) EXPECT() *MockPriceEventListenerMockRecorder {
	return m.recorder
}

// OnPriceEvent mocks base method.
func (m *MockPriceEventListener) OnPriceEvent(event *domain.PriceEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnPriceEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnPriceEvent indicates an expected call of OnPriceEvent.
func (mr *MockPriceEventListenerMockRecorder) OnPriceEvent(event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPriceEvent", reflect.TypeOf((*MockPriceEventListener)(nil).OnPriceEvent), event)
}

// MockNotifier is a mock of Notifier interface.
type MockNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierMockRecorder
	isgomock struct{}
}

// MockNotifierMockRecorder is the mock recorder for MockNotifier.
type MockNotifierMockRecorder struct {
	mock *MockNotifier
}

// NewMockNotifier creates a new mock instance.
func NewMockNotifier(ctrl *gomock.Controller) *MockNotifier {
	mock := &MockNotifier{ctrl: ctrl}
	mock.recorder = &MockNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifier) EXPECT() *MockNotifierMockRecorder {
	return m.recorder
}

// AddClient mocks base method.
func (m *MockNotifier) AddClient(ws *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddClient", ws)
}

// AddClient indicates an expected call of AddClient.
func (mr *MockNotifierMockRecorder) AddClient(ws any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockNotifier)(nil).AddClient), ws)
}

// Broadcast mocks base method.
func (m *MockNotifier) Broadcast(event *domain.PriceEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockNotifierMockRecorder) Broadcast(event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockNotifier)(nil).Broadcast), event)
}

// RemoveClient mocks base method.
func (m *MockNotifier) RemoveClient(ws *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveClient", ws)
}

// RemoveClient indicates an expected call of RemoveClient.
func (mr *MockNotifierMockRecorder) RemoveClient(ws any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveClient", reflect.TypeOf((*MockNotifier)(nil).RemoveClient), ws)
}

// Subscribe mocks base method.
func (m *MockNotifier) Subscribe(ws *websocket.Conn, stock domain.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ws, stock)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockNotifierMockRecorder) Subscribe(ws, stock any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockNotifier)(nil).Subscribe), ws, stock)
}

// Unsubscribe mocks base method.
func (m *MockNotifier) Unsubscribe(ws *websocket.Conn, stock domain.Stock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ws, stock)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockNotifierMockRecorder) Unsubscribe(ws, stock any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockNotifier)(nil).Unsubscribe), ws, stock)
}
