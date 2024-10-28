package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	slogmulti "github.com/samber/slog-multi"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"github.com/urodstvo/book-shop/libs/config"
)

type Logger interface {
	Info(input string, fields ...any)
	Error(input string, fields ...any)
	Debug(input string, fields ...any)
	Warn(input string, fields ...any)
	WithComponent(name string) Logger
	GetSlog() *slog.Logger
}

type Log struct {
	log *slog.Logger
}

type Opts struct {
	Env   string
	Level slog.Level
}

func NewFx() func(config config.Config) Logger {
	return func(config config.Config) Logger {
		return New(
			Opts{
				Env:   config.AppEnv,
				Level: slog.LevelInfo,
			},
		)
	}
}

func New(opts Opts) Logger {
	level := opts.Level

	var zeroLogWriter io.Writer
	if opts.Env == "production" {
		zeroLogWriter = os.Stderr
	} else {
		zeroLogWriter = zerolog.ConsoleWriter{Out: os.Stderr}
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	slogzerolog.SourceKey = "source"
	slogzerolog.ErrorKeys = []string{"error", "err"}
	zerolog.ErrorStackFieldName = "stack"

	zeroLogLogger := zerolog.New(zeroLogWriter)

	log := slog.New(
		slogmulti.Fanout(
			slogzerolog.Option{
				Level:     level,
				Logger:    &zeroLogLogger,
				AddSource: true,
			}.NewZerologHandler(),
		),
	)

	return &Log{
		log: log,
	}
}

func (c *Log) handle(level slog.Level, input string, fields ...any) {
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])
	r := slog.NewRecord(time.Now(), level, input, pcs[0])
	for _, f := range fields {
		r.Add(f)
	}
	_ = c.log.Handler().Handle(context.Background(), r)
}

func (c *Log) Info(input string, fields ...any) {
	c.handle(slog.LevelInfo, input, fields...)
}

func (c *Log) Warn(input string, fields ...any) {
	c.handle(slog.LevelWarn, input, fields...)
}

func (c *Log) Error(input string, fields ...any) {
	c.handle(slog.LevelError, input, fields...)

}

func (c *Log) Debug(input string, fields ...any) {
	c.handle(slog.LevelDebug, input, fields...)
}

func (c *Log) WithComponent(name string) Logger {
	return &Log{
		log: c.log.With(slog.String("component", name)),
	}
}

func (c *Log) GetSlog() *slog.Logger {
	return c.log
}
