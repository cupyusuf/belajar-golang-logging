package belajar_golang_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logger.Out = file

	logger.Info("Test Info")
	logger.Warn("Test Warn")
	logger.Error("Test Error")
}
