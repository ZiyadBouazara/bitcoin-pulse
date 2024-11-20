package domain

import (
	"encoding/json"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/models"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/ports"
	"github.com/gorilla/websocket"
	"sync"
)

type Server struct {
	conns         map[*websocket.Conn]bool
	subscriptions map[models.Stock]map[*websocket.Conn]bool
	logger        ports.Logger
	mu            sync.Mutex
}

func NewServer(logger ports.Logger) *Server {
	return &Server{
		conns:         make(map[*websocket.Conn]bool),
		subscriptions: make(map[models.Stock]map[*websocket.Conn]bool),
		logger:        logger,
	}
}

func (s *Server) HandleWS(ws *websocket.Conn) {
	s.logger.Info("new incoming connection from client: ", ws.RemoteAddr())

	s.mu.Lock()
	s.conns[ws] = true
	s.mu.Unlock()

	s.readLoop(ws)
}

func (s *Server) BroadcastPriceEvent(event *models.PriceEvent) {
	s.mu.Lock()
	clients := s.subscriptions[event.ProductID]
	defer s.mu.Unlock()

	if clients == nil {
		return
	}

	msg, err := json.Marshal(event)
	if err != nil {
		s.logger.Errorf("Error marshalling price event: %v", err)
		return
	}

	for ws := range clients {
		err := ws.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			s.logger.Errorf("Error sending message to client %v: %v", ws.RemoteAddr(), err)
			s.removeClient(ws)
			if err := ws.Close(); err != nil {
				return
			}
		}
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	defer func() {
		s.removeClient(ws)
		if err := ws.Close(); err != nil {
			return
		}
	}()
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			s.logger.Errorf("Error reading message: %v", err)
			break
		}

		var subMsg models.SubscriptionMessage
		if err := json.Unmarshal(message, &subMsg); err != nil {
			s.logger.Errorf("Invalid message format: %v", err)
			s.sendError(ws, "Invalid message format.")
			continue
		}

		if !models.IsSupportedStock(string(subMsg.Stock)) {
			s.logger.Errorf("Unsupported stock subscription attempt %v from client %v", subMsg.Stock, ws.RemoteAddr())
			s.sendError(ws, "Unsupported stock symbol")
			continue
		}

		switch subMsg.Action {
		case models.Subscribe:
			s.subscribeClient(ws, subMsg.Stock)
		case models.Unsubscribe:
			s.unsubscribeClient(ws, subMsg.Stock)
		default:
			s.logger.Errorf("Unknown action: %s", subMsg.Action)
			s.sendError(ws, "Unknown action")
		}

	}
}

func (s *Server) removeClient(ws *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.conns, ws)

	for stock, clients := range s.subscriptions {
		if clients[ws] {
			delete(clients, ws)
			if len(clients) == 0 {
				delete(s.subscriptions, stock)
			}
		}
	}
}

func (s *Server) subscribeClient(ws *websocket.Conn, stock models.Stock) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.subscriptions[stock] == nil {
		s.subscriptions[stock] = make(map[*websocket.Conn]bool)
	}
	s.subscriptions[stock][ws] = true
	s.logger.Infof("Client %v subscribed to %s", ws.RemoteAddr(), stock)
}

func (s *Server) unsubscribeClient(ws *websocket.Conn, stock models.Stock) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if clients, ok := s.subscriptions[stock]; ok {
		delete(clients, ws)
		if len(clients) == 0 {
			delete(s.subscriptions, stock)
		}
		s.logger.Infof("Client %v unsubscribed from %s", ws.RemoteAddr(), stock)
	}
}

func (s *Server) sendError(ws *websocket.Conn, errorMessage string) {
	errMsg := models.ErrorMessage{
		Type:    "error",
		Message: errorMessage,
	}

	message, err := json.Marshal(errMsg)
	if err != nil {
		s.logger.Errorf("Error marshalling error message: %v", err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		s.logger.Errorf("Error sending error message to client %v: %v", ws.RemoteAddr(), err)
	}
}
