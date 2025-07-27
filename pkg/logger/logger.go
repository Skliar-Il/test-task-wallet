package logger

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"github.com/Skliar-Il/test-task-wallet/pkg/exception"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

type KeyLoggerType string

const (
	RequestId KeyLoggerType = "request_id"
	lKey      KeyLoggerType = "middlewareLogger"
)

var middlewareLogger *Logger

type Config struct {
	Mode     string      `env:"LOGGER_MOD" default:"debug"`
	Topic    string      `env:"LOGGER_KAFKA_TOPIC" default:"log"`
	Name     string      `env:"LOGGER_NAME" default:"middlewareLogger"`
	KafkaCfg KafkaConfig `env:"KAFKA"`
}

type Logger struct {
	l *zap.Logger
}

func New(name string, level zapcore.LevelEnabler, broker sarama.SyncProducer, topicName string) *Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	kafkaCore := NewKafkaCore(broker, topicName, level, encoder)
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		level,
	)

	core := zapcore.NewTee(kafkaCore, consoleCore)
	localLogger := zap.New(core, zap.AddCaller())
	localLogger.Named(name)

	return &Logger{l: localLogger}
}

func GetLoggerFromCtx(ctx context.Context) *Logger {
	logger, exist := ctx.Value(lKey).(*Logger)
	if !exist {
		return nil
	}
	return logger
}

func (l *Logger) Sync() error {
	return l.l.Sync()
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

func SyncMiddleware() error {
	return middlewareLogger.Sync()
}

func Middleware(cfg *Config, broker sarama.SyncProducer) fiber.Handler {
	if cfg.Mode != "debug" && cfg.Mode != "production" {
		log.Fatalf("invalid middlewareLogger mod")
	}
	return func(c fiber.Ctx) error {
		guid := uuid.New().String()
		ctx := context.WithValue(c.Context(), RequestId, guid)

		if GetLoggerFromCtx(ctx) == nil {
			if middlewareLogger == nil {
				var err error
				if cfg.Mode == "debug" {
					middlewareLogger = New(cfg.Name, zapcore.DebugLevel, broker, cfg.Topic)
				} else {
					middlewareLogger = New(cfg.Name, zapcore.InfoLevel, broker, cfg.Topic)
				}
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString("middlewareLogger initialization failed")
				}
			}

			ctx = context.WithValue(ctx, lKey, middlewareLogger)
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

		code := c.Response().StatusCode()
		var fiberErr *fiber.Error
		var appErr *exception.AppException
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
		} else if errors.As(err, &appErr) {
			code = appErr.Code
		}

		if cfg.Mode == "debug" {
			GetLoggerFromCtx(ctx).Info(ctx,
				"Response HTTP",
				zap.Int("status", code),
				zap.String("response", c.Response().String()),
				zap.Duration("duration", duration),
			)
		} else if cfg.Mode == "production" {
			GetLoggerFromCtx(ctx).Info(ctx,
				"Response HTTP",
				zap.Int("status", code),
				zap.Duration("duration", duration),
			)
		}

		return err
	}
}
