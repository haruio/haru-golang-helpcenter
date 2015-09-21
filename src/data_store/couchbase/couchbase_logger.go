package couchbase

import (
	"log"

	"github.com/couchbaselabs/gocb/gocbcore"
)

type log_Logger struct {
}

var (
	globalLogger log_Logger
)

func (_ *log_Logger) Output(s string) error {
	log.Println(s)
	return nil
}

func LogStdOutLogger() gocbcore.Logger {
	return &globalLogger
}
