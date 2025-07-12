package logger

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"log"
	"time"
)

type KeyLoggerType string

const (
	RequestId KeyLoggerType = "request_id"
	lKey      KeyLoggerType = "logger"
)

type Config struct {
	Mode string `env:"LOGGER_MOD" default:"debug"`
}

type Logger struct {
	l *zap.Logger
}

func New(ctx context.Context) (context.Context, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, lKey, &Logger{logger})

	return ctx, nil
}

func GetLoggerFromCtx(ctx context.Context) *Logger {
	logger, exist := ctx.Value(lKey).(*Logger)
	if !exist {
		return nil
	}
	return logger
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(string(RequestId), ctx.Value(RequestId).(string)))
	}
	l.l.Info(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(string(RequestId), ctx.Value(RequestId).(string)))
	}
	l.l.Error(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(string(RequestId), ctx.Value(RequestId).(string)))
	}
	l.l.Fatal(msg, fields...)
}

func Middleware(cfg *Config) fiber.Handler {
	if cfg.Mode != "debug" && cfg.Mode != "production" {
		log.Fatalf("invalid logger mod")
	}
	return func(c fiber.Ctx) error {
		guid := uuid.New().String()
		ctx := context.WithValue(c.Context(), RequestId, guid)

		if GetLoggerFromCtx(ctx) == nil {
			var err error
			ctx, err = New(ctx)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("logger initialization failed")
			}
		}

		c.SetContext(ctx)

		if cfg.Mode == "debug" {
			GetLoggerFromCtx(ctx).Info(ctx,
				"Request HTTP",
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.ByteString("body", c.Body()),
			)
		} else if cfg.Mode == "production" {
			GetLoggerFromCtx(ctx).Info(ctx,
				"Request HTTP",
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
			)
		}

		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		if cfg.Mode == "debug" {
			GetLoggerFromCtx(ctx).Info(ctx,
				"Response HTTP",
				zap.Int("status", c.Response().StatusCode()),
				zap.String("response", c.Response().String()),
				zap.Duration("duration", duration),
			)
		} else if cfg.Mode == "production" {
			GetLoggerFromCtx(ctx).Info(ctx,
				"Response HTTP",
				zap.Int("status", c.Response().StatusCode()),
				zap.Duration("duration", duration),
			)
		}

		return err
	}
}
