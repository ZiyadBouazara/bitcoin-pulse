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

	return &LogrusLogger{logger}
}
