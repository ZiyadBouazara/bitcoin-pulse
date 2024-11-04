package cmd

import (
	"context"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain/services"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/infrastructure/kafka"
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/infrastructure/logging"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

var (
	priceService         *services.PriceService
	bitcoinPriceConsumer *kafka.BitcoinPriceConsumer
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	kafkaBrokerURL := os.Getenv("KAFKA_BROKER_URL")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaGroupID := os.Getenv("KAFKA_GROUP_ID")

	logger := logging.NewLogger()

	bitcoinPriceConsumer = kafka.NewBitcoinPriceConsumer(kafkaBrokerURL, kafkaTopic, kafkaGroupID, logger)
	priceService = services.NewPriceService(bitcoinPriceConsumer, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
		<-sigchan
		logger.Info("Received shutdown signal")
		cancel()
	}()

	// Start consuming
	priceService.StartConsuming(ctx)
}
