package logger

import "github.com/sirupsen/logrus"

func New(deployment string) *logrus.Logger {
	logger := logrus.New()

	switch deployment {
	case "local":
		logger.SetLevel(logrus.DebugLevel)
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	case "production":
		fallthrough
	default:
		logger.SetLevel(logrus.InfoLevel)
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	return logger
}
