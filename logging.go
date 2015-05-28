package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func BootstrapLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
}