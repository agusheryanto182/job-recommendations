package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// LogConfig holds configuration for logging
type LogConfig struct {
	// (debug, info, warn, error, fatal, panic)
	Level string

	// Output is the path to the log file
	Output string

	// format is the log format
	Format string
}

func NewLogger(cfg LogConfig) (*logrus.Logger, error) {
	logger := logrus.New()

	// Set level
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %v", err)
	}
	logger.SetLevel(level)

	logDir := filepath.Dir(cfg.Output)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	logFile, err := os.OpenFile(cfg.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	// Set output
	logger.SetOutput(logFile)

	// Set report caller
	logger.SetReportCaller(true)

	mw := io.MultiWriter(os.Stdout, logFile)
	logger.SetOutput(mw)

	// Set formatter
	switch cfg.Format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}

	return logger, nil
}

func InitLogger(config *Config) *logrus.Logger {
	cfg := LogConfig{
		Level:  "debug",
		Output: "log/app.log",
		Format: "json",
	}

	if config != nil {
		if config.LogLevel != "" {
			cfg.Level = config.LogLevel
		}
		if config.LogFile != "" {
			cfg.Output = config.LogFile
		}
		if config.LogFormat != "" {
			cfg.Format = config.LogFormat
		}
	}

	if level := os.Getenv("LOG_LEVEL"); level != "" {
		cfg.Level = level
	}
	if output := os.Getenv("LOG_FILE"); output != "" {
		cfg.Output = output
	}
	if format := os.Getenv("LOG_FORMAT"); format != "" {
		cfg.Format = format
	}

	logger, err := NewLogger(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize logger: %v", err))
	}

	logrus.SetOutput(logger.Out)
	logrus.SetFormatter(logger.Formatter)
	logrus.SetLevel(logger.Level)
	logrus.SetReportCaller(true)

	return logger
}
