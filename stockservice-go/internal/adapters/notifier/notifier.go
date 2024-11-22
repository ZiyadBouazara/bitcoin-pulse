package notifier

import (
	"encoding/json"
	"fmt"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/ports"
	"github.com/gorilla/websocket"
	"sync"
)

type Notifier struct {
	conns         map[*websocket.Conn]bool
	subscriptions map[domain.Stock]map[*websocket.Conn]bool
	logger        ports.Logger
	mu            sync.Mutex
}

func NewNotifier(logger ports.Logger) *Notifier {
	return &Notifier{
		conns:         make(map[*websocket.Conn]bool),
		subscriptions: make(map[domain.Stock]map[*websocket.Conn]bool),
		logger:        logger,
		mu:            sync.Mutex{},
	}
}

func (n *Notifier) AddClient(ws *websocket.Conn) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.conns[ws] = true
}

func (n *Notifier) RemoveClient(ws *websocket.Conn) {
	n.mu.Lock()
	defer n.mu.Unlock()

	delete(n.conns, ws)

	for stock, clients := range n.subscriptions {
		if clients[ws] {
			delete(clients, ws)
			if len(clients) == 0 {
				delete(n.subscriptions, stock)
			}
		}
	}
}

func (n *Notifier) Broadcast(event *domain.PriceEvent) error {
	if event == nil {
		return fmt.Errorf("received a nil PriceEvent")
	}

	n.mu.Lock()
	clients := n.subscriptions[event.ProductID]
	defer n.mu.Unlock()

	if clients == nil {
		return nil
	}

	msg, err := json.Marshal(event)
	if err != nil {
		n.logger.Errorf("Error marshalling price event: %v", err)
		return nil
	}

	for ws := range clients {
		err := ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			n.logger.Errorf("Error sending message to client %v: %v", ws.RemoteAddr(), err)
			n.RemoveClient(ws)
			if err := ws.Close(); err != nil {
				return nil
			}
		}
	}

	return nil
}

func (n *Notifier) Subscribe(ws *websocket.Conn, stock domain.Stock) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.subscriptions[stock] == nil {
		n.subscriptions[stock] = make(map[*websocket.Conn]bool)
	}
	n.subscriptions[stock][ws] = true
	n.logger.Infof("Client %v subscribed to %v", ws.RemoteAddr(), stock)

	return nil
}

func (n *Notifier) Unsubscribe(ws *websocket.Conn, stock domain.Stock) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if clients, ok := n.subscriptions[stock]; ok {
		delete(clients, ws)
		if len(clients) == 0 {
			delete(n.subscriptions, stock)
		}
		n.logger.Infof("Client %v unsubscribed from %v", ws.RemoteAddr(), stock)
	}

	return nil
}
