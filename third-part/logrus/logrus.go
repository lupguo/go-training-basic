package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

var Environment = "dev"

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: false,
	})
	switch Environment {
	case "dev", "test":
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetOutput(os.Stdout)
	case "prod":
		logrus.SetLevel(logrus.InfoLevel)
		logFile := os.TempDir() + "logrus.log"
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.SetOutput(file)
	}
}

func main() {
	// log to file
	FileLogging()

	// Trace、Debug、Info、Warn、Error、Fatal、Panic
	DefaultLogging()

	// log Context (common part)
	CommonLogging()

	// log http server message
	ServerLogging()
}

func ServerLogging() {
	w := logrus.StandardLogger().Writer()
	_ = &http.Server{
		ErrorLog: log.New(w, "", 0),
	}
}

func CommonLogging() {
	hostname, _ := os.Hostname()

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logrus.WithFields(logrus.Fields{
		"hostname": hostname,
		"pid":      os.Getpid(),
	})
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}

func DefaultLogging() {
	dfLogger := logrus.WithFields(logrus.Fields{
		"uid":      100,
		"nickname": "clark",
		"new":      false,
	})

	dfLogger.Trace("---------trace level msg")
	dfLogger.Debug("---------debug level msg")
	dfLogger.Info("---------info level msg")
	dfLogger.Warn("---------warn level msg")
	dfLogger.Error("---------error level msg")
	//dfLogger.Fatal("user join activity !")
	//dfLogger.Panic("panic level msg @")
}

func FileLogging() {
	fileLogger := logrus.New()
	logFile := os.TempDir() + "logrus.log"

	//You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		fileLogger.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}

	fileLogger.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warn("A group of walrus emerges from the ocean")

	// log output os.Stdout
	logrus.WithFields(logrus.Fields{
		"logFile": logFile,
		"version": "v1.0.1",
	}).Info("logrus to file...")
}
