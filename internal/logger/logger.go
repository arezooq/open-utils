package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	serviceName string
	logger      *log.Logger
}

func New(serviceName string) *Logger {
	return &Logger{
		serviceName: serviceName,
		logger:      log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) log(level LogLevel, requestID, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	entry := fmt.Sprintf("[%s] [%s] [%s] [%s] %s",
		timestamp, l.serviceName, requestID, level, message)

	switch level {
	case ERROR, WARN, INFO, DEBUG:
		l.logger.Println(entry)
	case FATAL:
		l.logger.Println(entry)
		os.Exit(1)
	case PANIC:
		l.logger.Panicln(entry)
	}
}

func (l *Logger) Info(reqID, msg string)  { l.log(INFO, reqID, msg) }
func (l *Logger) Warn(reqID, msg string)  { l.log(WARN, reqID, msg) }
func (l *Logger) Error(reqID, msg string) { l.log(ERROR, reqID, msg) }
func (l *Logger) Debug(reqID, msg string) { l.log(DEBUG, reqID, msg) }
func (l *Logger) Fatal(reqID, msg string) { l.log(FATAL, reqID, msg) }
func (l *Logger) Panic(reqID, msg string) { l.log(PANIC, reqID, msg) }
