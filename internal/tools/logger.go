package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// LogLevel represents different levels of logging severity.
type LogLevel int8

const (
	INFO  LogLevel = iota // Informational messages
	DEBUG                 // Debug-level messages
	WARN                  // Warnings that require attention
	ERROR                 // Errors that need immediate fixing
)

var log zerolog.Logger
var logFuncs map[LogLevel]func(interface{})

// Log writes a log message at the specified level.
// If an unknown LogLevel is provided, it defaults to INFO.
//
// Example:
//
//	logger.Log("Service started", logger.INFO)
//	logger.Log("Debugging details", logger.DEBUG)
func Log(message interface{}, level LogLevel) {
	if logFunc, exists := logFuncs[level]; exists {
		logFunc(message)
	} else {
		Log(message, INFO) // Default to INFO if level is invalid
	}
}

func init() {
	debugMode := strings.EqualFold(os.Getenv("DEBUG"), "true")
	logLevel := getLogLevel(debugMode)
	output := getOutputWriter(debugMode)

	log = zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + 2).Logger().Level(logLevel)

	// Function map to associate LogLevel with the correct logging function.
	logFuncs = map[LogLevel]func(interface{}){
		INFO:  func(m interface{}) { log.Info().Msg(fmt.Sprint(m)) },
		DEBUG: func(m interface{}) { log.Debug().Msg(fmt.Sprint(m)) },
		WARN:  func(m interface{}) { log.Warn().Msg(fmt.Sprint(m)) },
		ERROR: func(m interface{}) { log.Error().Msg(fmt.Sprint(m)) },
	}
}

// getLogLevel returns the appropriate zerolog.Level based on the debug mode.
// If DEBUG is enabled, the log level is set to DebugLevel; otherwise, InfoLevel.
func getLogLevel(debug bool) zerolog.Level {
	if debug {
		return zerolog.DebugLevel
	}
	return zerolog.InfoLevel
}

// getOutputWriter determines the log output format.
// If debug mode is enabled, it uses a human-readable console output.
// Otherwise, logs are formatted as structured JSON.
func getOutputWriter(debug bool) io.Writer {
	if debug {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		consoleWriter.NoColor = true

		consoleWriter.FormatLevel = func(i interface{}) string { return fmt.Sprintf("| %-6s |", strings.ToUpper(fmt.Sprint(i))) }
		consoleWriter.FormatMessage = func(i interface{}) string { return fmt.Sprintf("%s", i) }
		consoleWriter.FormatFieldName = func(i interface{}) string { return fmt.Sprintf("%s:", i) }
		consoleWriter.FormatFieldValue = func(i interface{}) string { return fmt.Sprintf("%v", i) }

		return consoleWriter
	}
	return os.Stderr
}
