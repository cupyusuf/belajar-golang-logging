package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "yusuf",
		"message":  "ini adalah pesan",
	}).Info("ini adalah log")

	logger.WithField("username", "yusuf").WithField("name", "yusuf").Info("ini adalah log")
}
