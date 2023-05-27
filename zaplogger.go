package redis

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() (*ZapLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{
		logger: logger,
	}, nil
}

func (z *ZapLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	z.logger.Info(fmt.Sprintf(format, v...))
}

func (z *ZapLogger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Debug(fmt.Sprintf(msg, keysAndValues...))
}

func (z *ZapLogger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Info(fmt.Sprintf(msg, keysAndValues...))
}

func (z *ZapLogger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Warn(fmt.Sprintf(msg, keysAndValues...))
}

func (z *ZapLogger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Error(fmt.Sprintf(msg, keysAndValues...))
}

func (z *ZapLogger) Fatal(ctx context.Context, msg string, keysAndValues ...interface{}) {
	z.logger.Fatal(fmt.Sprintf(msg, keysAndValues...))
}
