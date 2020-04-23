package mocks

import (
	"os"

	"github.com/sirupsen/logrus"
)

func TestLogger() *logrus.Entry {
	log := logrus.New()
	log.Out = os.Stderr
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "02/Jan/2006:15:04:05",
	}
	log.Level = logrus.PanicLevel
	return logrus.NewEntry(log)
}
