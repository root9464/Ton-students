package utils

import (
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "15:04:05",
		FullTimestamp:   true,
	})

	return log
}
