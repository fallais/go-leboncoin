package main

import (
	"flag"
	"net"
	"net/http"

	"go-leboncoin/runner"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/zenazn/goji/graceful"
)

var (
	bindAddress = flag.String("bind", ":8000", "Network address used to bind")
	logging      = flag.String("logging", "info", "Logging level")
	filtersFile  = flag.String("filters_file", "filters.yaml", "Filters")
	runOnStartup = flag.Bool("run_on_startup", true, "Run on startup ?")
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
	// Create the runner
	r := runner.New(*filtersFile)

	if *runOnStartup {
		r.Run()
	}

	// Initialize CRON
	c := cron.New()
	c.AddFunc("@every 5m", r.Run)
	c.Start()

	// Handlers
	http.Handle("/metrics", promhttp.Handler())

	// Initialize the goroutine listening to signals passed to the app
	graceful.HandleSignals()

	// Pre-graceful shutdown event
	graceful.PreHook(func() {
		logrus.Infoln("Received a signal, stopping the application")
	})

	// Post-shutdown event
	graceful.PostHook(func() {
		// Stop all the taks
		c.Stop()

		logrus.Infoln("Stopped the application")
	})

	// Listen to the passed address
	logrus.Infoln("Starting the Web server")
	listener, err := net.Listen("tcp", *bindAddress)
	if err != nil {
		logrus.Fatalln("Cannot set up a TCP listener")
	}
	logrus.Infoln("Successfully started the Web server")

	// Start the listening
	err = graceful.Serve(listener, http.DefaultServeMux)
	if err != nil {
		logrus.Errorf("Error with the server : %s", err)
	}

	// Wait until open connections close
	graceful.Wait()
}
