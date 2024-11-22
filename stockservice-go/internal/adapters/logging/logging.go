package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogger() *LogrusLogger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logger.SetLevel(logrus.InfoLevel)

	return &LogrusLogger{logger}
}
