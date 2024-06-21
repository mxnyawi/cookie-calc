package logger

import (
	"log"
	"os"
)

// LoggerInterface defines the logging methods that your application uses.
type LoggerInterface interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	Close()
}

type Logger struct {
	File *log.Logger
}

func SetupLogger() LoggerInterface {
	// Open a file for logging
	logFile, err := os.OpenFile("cookie_calc.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	// Logger for file output
	fileLogger := log.New(logFile, "FILE: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		File: fileLogger,
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.File.Printf("INFO: "+msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.File.Printf("ERROR: "+msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.File.Printf("FATAL: "+msg, args...)
}

func (l *Logger) Close() {
	l.File.Writer().(*os.File).Close()
}

func StdOut() *log.Logger {
	return log.New(os.Stdout, "STDOUT: ", log.Ldate|log.Ltime|log.Lshortfile)
}

type NoOpLogger struct{}

func (n *NoOpLogger) Info(msg string, args ...interface{}) {
	// Do nothing
}

func (n *NoOpLogger) Error(msg string, args ...interface{}) {
	// Do nothing
}

func (n *NoOpLogger) Fatal(msg string, args ...interface{}) {
	// Do nothing
}

func (n *NoOpLogger) Close() {
	// Do nothing
}

// NewNoOpLogger creates and returns a new NoOpLogger instance
func NewNoOpLogger() LoggerInterface {
	return &NoOpLogger{}
}
