package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logger.WithFields(logrus.Fields{
		"username": "yusuf",
		"age":      25,
	})

	entry.Info("ini adalah log info")
}
