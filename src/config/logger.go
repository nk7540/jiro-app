package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Logger() {
	format := &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}

	logrus.SetFormatter(format)
	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}
