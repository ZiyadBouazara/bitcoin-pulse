package main

import (
	"context"
	"errors"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/config"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/adapters/handlers"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/adapters/kafka"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/adapters/logging"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/adapters/notifier"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

var (
	cfg               *config.Config
	router            *gin.Engine
	logger            *logging.LogrusLogger
	priceService      *services.PriceService
	livePricesHandler *handlers.LivePricesHandler
	notif             *notifier.Notifier
)

func main() {
	cfg = loadConfig()
	logger = logging.NewLogger()
	notif = notifier.NewNotifier(logger)

	priceService = initKafkaConsumer()

	router = initRoutes()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := startHTTPServer()

	go handleShutdown(cancel, srv)

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

func initRoutes() *gin.Engine {
	router := gin.Default()

	livePricesHandler = handlers.NewLivePricesHandler(priceService, logger)
	router.GET("/ws/livepricesfeed", livePricesHandler.HandleWebSocket)

	return router
}

func initKafkaConsumer() *services.PriceService {
	bitcoinPriceConsumer := kafka.NewBitcoinPriceConsumer(
		cfg.KafkaBrokerURL,
		cfg.KafkaTopic,
		cfg.KafkaGroupID,
		logger,
	)
	priceService = services.NewPriceService(notif, bitcoinPriceConsumer, logger)
	return priceService
}

func startHTTPServer() *http.Server {
	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: router,
	}
	go func() {
		logger.Infof("Server started on %v", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("Server failed: %v", err)
		}
	}()
	return srv
}

func handleShutdown(cancel context.CancelFunc, srv *http.Server) {
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
