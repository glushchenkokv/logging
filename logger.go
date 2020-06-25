package logging

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// Config configures Logger
type Config struct {
	Level            string `yaml:"level"`                               // The logging level the logger should log at (`error` by default).
	Format           string `yaml:"format"`                              // The logging format (`text` or `json`) (`json` by default).
	ReportCaller     bool   `yaml:"reportCaller"     split_words:"true"` // Flag for whether to log caller info (off by default).
	DisableTimestamp bool   `yaml:"disableTimestamp" split_words:"true"` // DisableTimestamp allows disabling automatic timestamps in output (off by default).
}

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(cfg Config) *Logger {
	return &Logger{
		logger: &logrus.Logger{
			Out:          os.Stdout,
			Formatter:    newFormatter(cfg.Format, cfg.DisableTimestamp),
			ReportCaller: cfg.ReportCaller,
			Level:        getLevel(cfg.Level),
		},
	}
}

func getLevel(level string) logrus.Level {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return logrus.ErrorLevel
	}

	return lvl
}

func newFormatter(format string, disableTimestamp bool) logrus.Formatter {
	timestampFormat := "2006-01-02 15:04:05.000"

	jsonFmt := &logrus.JSONFormatter{
		DisableTimestamp: disableTimestamp,
		TimestampFormat:  timestampFormat,
	}

	switch strings.ToLower(strings.TrimSpace(format)) {
	case "text":
		return &logrus.TextFormatter{
			DisableTimestamp: disableTimestamp,
			TimestampFormat:  timestampFormat,
			FullTimestamp:    true,
			ForceColors:      true,
		}
	case "json":
		return jsonFmt
	default:
		return jsonFmt
	}
}
