package logging

import (
	"github.com/ZiyadBouazara/bitcoin-pulse/stockservice-go/internal/domain/ports"
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogger() ports.Logger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &LogrusLogger{logger}
}

//var _ ports.Logger = (*LogrusLogger)(nil)
