package main

import (
	"context"
	"errors"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/config"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/domain"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/ports"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/infrastructure/kafka"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/infrastructure/logging"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

func main() {
	cfg := loadConfig()
	logger := logging.NewLogger()
	server := domain.NewServer(logger)
	priceService := initKafkaConsumer(cfg, server, logger)

	initRoutes(server, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := startHTTPServer(cfg.Port, logger)

	go handleShutdown(cancel, logger, srv)

	priceService.StartConsuming(ctx)
}

func loadConfig() *config.Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":3000"
	}

	kafkaBrokerURL := os.Getenv("KAFKA_BROKER_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaGroupID := os.Getenv("KAFKA_GROUP_ID")

	if kafkaBrokerURL == "" || kafkaTopic == "" || kafkaGroupID == "" {
		panic("Kafka configuration environment variables are not set.")
	}

	return &config.Config{
		Port:           port,
		KafkaBrokerURL: kafkaBrokerURL,
		KafkaTopic:     kafkaTopic,
		KafkaGroupID:   kafkaGroupID,
	}
}

func initRoutes(server *domain.Server, logger ports.Logger) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// TODO: Secure this in production by checking the Origin header
			return true
		},
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Errorf("Failed to upgrade to WebSocket: %v", err)
			return
		}

		go server.HandleWS(ws)
	})
}

func initKafkaConsumer(config *config.Config, server ports.Server, logger ports.Logger) *services.PriceService {
	bitcoinPriceConsumer := kafka.NewBitcoinPriceConsumer(
		config.KafkaBrokerURL,
		config.KafkaTopic,
		config.KafkaGroupID,
		logger,
	)
	priceService := services.NewPriceService(server, bitcoinPriceConsumer, logger)
	return priceService
}

func startHTTPServer(port string, logger ports.Logger) *http.Server {
	srv := &http.Server{Addr: port}
	go func() {
		logger.Infof("Server started on %v", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("Server failed: %v", err)
		}
	}()
	return srv
}

func handleShutdown(cancel context.CancelFunc, logger ports.Logger, srv *http.Server) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigchan
	logger.Infof("Received shutdown signal %v. Initiating shutdown... 👋", sig)

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Errorf("Server forced to shutdown: %v", err)
	}

	cancel()
}
