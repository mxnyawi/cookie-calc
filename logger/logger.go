package logger

import (
	"log"
	"os"
)

// LoggerInterface defines the logging methods
type LoggerInterface interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Close()
}

// Logger represents the logger instance
type Logger struct {
	File *log.Logger
}

// SetupLogger creates and returns a new Logger instance
func SetupLogger() LoggerInterface {
	logFile, err := os.OpenFile("cookie_calc.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	fileLogger := log.New(logFile, "FILE: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		File: fileLogger,
	}
}

// Info logs an info message
func (l *Logger) Info(msg string, args ...interface{}) {
	l.File.Printf("INFO: "+msg, args...)
}

// Error logs an error message
func (l *Logger) Error(msg string, args ...interface{}) {
	l.File.Printf("ERROR: "+msg, args...)
}

// Close closes the log file
func (l *Logger) Close() {
	l.File.Writer().(*os.File).Close()
}

// StdOut returns a new log.Logger instance that writes to os.Stdout
func StdOut() *log.Logger {
	return log.New(os.Stdout, "", 0)
}

// NoOpLogger is a no-op logger
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
