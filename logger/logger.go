package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// Set up the logger configuration
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout) // You can change this to a file later
	log.SetLevel(logrus.InfoLevel)
}

func Info(msg string, fields logrus.Fields) {
	log.WithFields(fields).Info(msg)
}

func Error(msg string, fields logrus.Fields) {
	log.WithFields(fields).Error(msg)
}

func Warn(msg string, fields logrus.Fields) {
	log.WithFields(fields).Warn(msg)
}

func Debug(msg string, fields logrus.Fields) {
	log.WithFields(fields).Debug(msg)
}
