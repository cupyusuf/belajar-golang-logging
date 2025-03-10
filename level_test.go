package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // set level

	logger.Trace("this is a trace")
	logger.Debug("this is a debug")
	logger.Info("this is a info")
	logger.Warn("this is a warn")
	logger.Error("this is a error")
	// logger.Fatal("this is a fatal")
	// logger.Panic("this is a panic")
}
