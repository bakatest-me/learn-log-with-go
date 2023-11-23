package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(level)
	return logger
}
func main() {
	level := logrus.InfoLevel
	log := NewLogger(level)

	log.Trace("Print Trace")
	log.Debug("Print Debug") //use this for debug program
	log.Info("Print Info")
	log.Warn("Print Warn")
	log.Error("Print Error")
	log.Fatal("Print Fatal") //FATAL log will stop program

	log.Info("working")
	log.Fatal("Stop") //FATAL log will stop program
	log.Info("working")

}
