package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
    serviceName string
    requestID   string
    logger      *log.Logger
}

func New(serviceName string) *Logger {
    return &Logger{
        serviceName: serviceName,
        logger:      log.New(os.Stdout, "", 0),
    }
}

func (l *Logger) WithRequestID(reqID string) *Logger {
    return &Logger{
        serviceName: l.serviceName,
        requestID:   reqID,
        logger:      l.logger,
    }
}

func (l *Logger) log(level LogLevel, message string) {
    timestamp := time.Now().Format(time.RFC3339)
    entry := fmt.Sprintf("[%s] [%s] [%s] [%s] %s",
        timestamp, l.serviceName, l.requestID, level, message)

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

func (l *Logger) Info(msg string)  { l.log(INFO, msg) }
func (l *Logger) Warn(msg string)  { l.log(WARN, msg) }
func (l *Logger) Error(msg string) { l.log(ERROR, msg) }
func (l *Logger) Debug(msg string) { l.log(DEBUG, msg) }
func (l *Logger) Fatal(msg string) { l.log(FATAL, msg) }
func (l *Logger) Panic(msg string) { l.log(PANIC, msg) }
