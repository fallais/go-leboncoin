package main

import (
	"flag"

	"go-leboncoin/runner"

	"github.com/sirupsen/logrus"
)

var (
	logging     = flag.String("logging", "info", "Logging level")
	filtersFile = flag.String("filters_file", "filters.yaml", "Filters")
)

func init() {
	// Parse the flags
	flag.Parse()

	// Set the logging level
	switch *logging {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// Set the TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	logrus.Infoln("go-leboncoin is starting")
}

func main() {
	runner.Run(*filtersFile)
}
