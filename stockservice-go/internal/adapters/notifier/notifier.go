package notifier

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/ports"
	"github.com/gorilla/websocket"
)

type Notifier struct {
	conns         sync.Map // key: ports.WebSocketConn, value: struct{}
	subscriptions sync.Map // key: domain.Stock, value: *sync.Map (key: ports.WebSocketConn, value: struct{})
	logger        ports.Logger
}

func NewNotifier(logger ports.Logger) *Notifier {
	return &Notifier{
		logger: logger,
	}
}

func (n *Notifier) AddClient(ws ports.WebSocketConn) {
	n.conns.Store(ws, struct{}{})
}

func (n *Notifier) RemoveClient(ws ports.WebSocketConn) {
	n.conns.Delete(ws)

	n.subscriptions.Range(func(key, value interface{}) bool {
		clients := value.(*sync.Map)
		clients.Delete(ws)
		return true
	})
}

func (n *Notifier) Subscribe(ws ports.WebSocketConn, stock domain.Stock) error {
	clientsInterface, _ := n.subscriptions.LoadOrStore(stock, &sync.Map{})
	clients := clientsInterface.(*sync.Map)
	clients.Store(ws, struct{}{})
	n.logger.Infof("Client %v subscribed to %v", ws.RemoteAddr(), stock)
	return nil
}

func (n *Notifier) Unsubscribe(ws ports.WebSocketConn, stock domain.Stock) error {
	clientsInterface, ok := n.subscriptions.Load(stock)
	if ok {
		clients := clientsInterface.(*sync.Map)
		clients.Delete(ws)
		n.logger.Infof("Client %v unsubscribed from %v", ws.RemoteAddr(), stock)
	}
	return nil
}

func (n *Notifier) Broadcast(event *domain.PriceEvent) error {
	if event == nil {
		return fmt.Errorf("received a nil PriceEvent")
	}

	clientsInterface, ok := n.subscriptions.Load(event.ProductID)
	if !ok {
		return nil
	}

	clients := clientsInterface.(*sync.Map)
	msg, err := json.Marshal(event)
	if err != nil {
		n.logger.Errorf("Error marshalling price event: %v", err)
		return nil
	}

	clients.Range(func(key, _ interface{}) bool {
		ws := key.(ports.WebSocketConn)
		if err := ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			n.logger.Errorf("Error sending message to client %v: %v", ws.RemoteAddr(), err)
			clients.Delete(ws)
			n.conns.Delete(ws)
			if err := ws.Close(); err != nil {
				n.logger.Errorf("Error closing WebSocket: %v", err)
			}
		}
		return true
	})

	return nil
}

func (n *Notifier) GetConnections() map[ports.WebSocketConn]struct{} {
	connsCopy := make(map[ports.WebSocketConn]struct{})
	n.conns.Range(func(key, value interface{}) bool {
		ws := key.(ports.WebSocketConn)
		connsCopy[ws] = struct{}{}
		return true
	})
	return connsCopy
}

func (n *Notifier) GetSubscriptions(stock domain.Stock) map[ports.WebSocketConn]struct{} {
	clientsCopy := make(map[ports.WebSocketConn]struct{})
	clientsInterface, ok := n.subscriptions.Load(stock)
	if !ok {
		return clientsCopy // Return empty map
	}

	clients := clientsInterface.(*sync.Map)
	clients.Range(func(key, _ interface{}) bool {
		ws := key.(ports.WebSocketConn)
		clientsCopy[ws] = struct{}{}
		return true
	})
	return clientsCopy
}
