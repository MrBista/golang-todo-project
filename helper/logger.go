package helper

import (
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	logger := logrus.New()

	level := logrus.InfoLevel

	logger.SetLevel(level)

	return logger
}
