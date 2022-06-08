package application

import (
	"log"
)

type Logger interface {
	Debug(s string, i ...interface{})
	Info(s string, i ...interface{})
	Warn(s string, i ...interface{})
	Error(s string, i ...interface{})
	Panic(s string, i ...interface{})
}

type DefaultLogger struct{}

func NewDefaultLogger() Logger {
	return &DefaultLogger{}
}

func (d *DefaultLogger) Debug(s string, i ...interface{}) {
	log.Println("DEBUG: ", s, i)
}

func (d *DefaultLogger) Info(s string, i ...interface{}) {
	log.Println("INFO: ", s, i)
}

func (d *DefaultLogger) Warn(s string, i ...interface{}) {
	log.Println("WARN: ", s, i)
}

func (d *DefaultLogger) Error(s string, i ...interface{}) {
	log.Fatalln("ERROR: ", s, i)
}

func (d *DefaultLogger) Panic(s string, i ...interface{}) {
	log.Fatalln("PANIC: ", s, i)
}
