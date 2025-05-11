package logger

import "github.com/sirupsen/logrus"

func New() *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	return logger
}
