package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormater(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.Warn("Hello Logging")
	logger.Error("Hello Logging")
}
