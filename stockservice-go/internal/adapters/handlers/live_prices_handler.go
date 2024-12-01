package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin:     func(r *http.Request) bool { return true }, // TODO: Secure this in production by checking the Origin header
}

type LivePricesHandler struct {
	priceService ports.PriceService
	logger       ports.Logger
}

func NewLivePricesHandler(ps ports.PriceService, logger ports.Logger) *LivePricesHandler {
	return &LivePricesHandler{
		priceService: ps,
		logger:       logger,
	}
}

func (h *LivePricesHandler) HandleWebSocket(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		h.logger.Errorf("Failed to upgrade to WebSocket: %v", err)
		return
	}

	go h.handleConnection(ctx, ws)
}

func (h *LivePricesHandler) handleConnection(ctx *gin.Context, ws ports.WebSocketConn) {
	h.logger.Infof("New client connected: %v", ws.RemoteAddr())
	defer func() {
		h.priceService.RemoveClient(ws)
		if err := ws.Close(); err != nil {
			h.logger.Errorf("Error closing WebSocket: %v", err)
		}
		h.logger.Infof("Client disconnected: %v", ws.RemoteAddr())
	}()

	h.priceService.AddClient(ws)

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			h.logger.Errorf("Error reading message: %v", err)
			break
		}

		var subMsg domain.SubscriptionMessage
		if err := json.Unmarshal(message, &subMsg); err != nil {
			h.logger.Errorf("Invalid message format: %v", err)
			h.sendError(ws, "Invalid message format.")
			continue
		}

		if !domain.IsSupportedStock(string(subMsg.Stock)) {
			h.logger.Errorf("Unsupported stock subscription attempt %v from client %v", subMsg.Stock, ws.RemoteAddr())
			h.sendError(ws, "Unsupported stock symbol")
			continue
		}

		switch subMsg.Action {
		case domain.Subscribe:
			err := h.priceService.Subscribe(ws, subMsg.Stock)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, fmt.Errorf("failed to subscribe client: %v", err))
				return
			}
		case domain.Unsubscribe:
			err := h.priceService.Unsubscribe(ws, subMsg.Stock)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, fmt.Errorf("failed to unsubscribe client: %v", err))
				return
			}
		default:
			h.logger.Errorf("Unknown action: %s", subMsg.Action)
			h.sendError(ws, "Unknown action")
		}
	}
}

func (h *LivePricesHandler) sendError(ws ports.WebSocketConn, errorMessage string) {
	errMsg := domain.ErrorMessage{
		Type:    "error",
		Message: errorMessage,
	}

	message, err := json.Marshal(errMsg)
	if err != nil {
		h.logger.Errorf("Error marshalling error message: %v", err)
		return
	}

	if err := ws.WriteMessage(websocket.TextMessage, message); err != nil {
		h.logger.Errorf("Error sending error message to client %v: %v", ws.RemoteAddr(), err)
	}
}
