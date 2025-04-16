package logger

import (
	"fmt"
	"log"
	"os"
)

type Attrs map[string]any

type Logger struct {
	logger *log.Logger
}

func New() *Logger {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) logMessage(level, msg string, attrs ...Attrs) {
	logMessage := fmt.Sprintf("[%s] %s", level, msg)

	if len(attrs) > 0 {
		for key, value := range attrs {
			logMessage += fmt.Sprintf(" %d=%v", key, value)
		}
	}

	l.logger.Println(logMessage)
}

func (l *Logger) Info(msg string, attrs ...Attrs) {
	l.logMessage("INFO", msg, attrs...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.logMessage("INFO", msg)
}

func (l *Logger) Warn(msg string, attrs ...Attrs) {
	l.logMessage("WARN", msg, attrs...)
}

func (l *Logger) Error(msg string, attrs ...Attrs) {
	l.logMessage("ERROR", msg, attrs...)
}

func (l *Logger) Errorf(msg string, params ...interface{}) {
	l.Error(fmt.Sprintf(msg, params...))
}

func (l *Logger) Fatal(msg string, attrs ...Attrs) {
	l.logMessage("FATAL", msg, attrs...)
	os.Exit(1)
}

func (l *Logger) Fatalf(msg string, params ...interface{}) {
	l.Fatal(fmt.Sprintf(msg, params...))
}

func (l *Logger) Debug(msg string, attrs ...Attrs) {
	l.logMessage("DEBUG", msg, attrs...)
}

func (l *Logger) Logf(level, format string, params ...interface{}) {
	l.logMessage(level, fmt.Sprintf(format, params...), nil)
}
