/*
Logger wrapper around the log/slog package to enable changing the log level
dynamically.
*/
package log

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

var (
	lvl    = new(slog.LevelVar)
	logger *slog.Logger
)

func init() {
	lvl.Set(slog.LevelInfo)

	logger = slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      lvl,
		TimeFormat: time.Kitchen,
	}))
}

func SetDebug() {
	lvl.Set(slog.LevelDebug)
}

func SetInfo() {
	lvl.Set(slog.LevelInfo)
}

func SetWarn() {
	lvl.Set(slog.LevelWarn)
}

func SetError() {
	lvl.Set(slog.LevelError)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
