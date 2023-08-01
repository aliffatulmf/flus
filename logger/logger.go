package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

func SetOutput(w io.Writer) {
	infoLogger.SetOutput(w)
	warnLogger.SetOutput(w)
	errorLogger.SetOutput(w)
}

func Info(args ...interface{}) {
	infoLogger.Output(2, fmt.Sprintln(args...))
}

func Infof(format string, args ...interface{}) {
	infoLogger.Output(2, fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	warnLogger.Output(2, fmt.Sprintln(args...))
}

func Warnf(format string, args ...interface{}) {
	warnLogger.Output(2, fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	errorLogger.Output(2, fmt.Sprintln(args...))
}

func Errorf(format string, args ...interface{}) {
	errorLogger.Output(2, fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	errorLogger.Output(2, fmt.Sprintln(args...))
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	errorLogger.Output(2, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Panic(args ...interface{}) {
	s := fmt.Sprintln(args...)
	errorLogger.Output(2, s)
	panic(strings.TrimSpace(s))
}

func Panicf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	errorLogger.Output(2, s)
	panic(strings.TrimSpace(s))
}
