package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)

	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	Log.SetFormatter(formatter)
}
