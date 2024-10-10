package logger

import (
	"github.com/lmittmann/tint"
	"io"
	"log/slog"
	"time"
)

type SlogLogger interface {
	Info(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

func MustTintSlogLogger(w io.Writer, opts *slog.HandlerOptions) SlogLogger {
	if opts == nil {
		opts = &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
	}

	return slog.New(tint.NewHandler(w, &tint.Options{
		AddSource:   opts.AddSource,
		Level:       opts.Level,
		ReplaceAttr: opts.ReplaceAttr,
		TimeFormat:  time.DateTime,
	}))
}
