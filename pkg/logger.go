package logger

import (
	"github.com/lhdhtrc/logger-go/pkg/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Console bool `json:"console" bson:"console" yaml:"console" mapstructure:"console"`
	Remote  bool `json:"remote" bson:"remote" yaml:"remote" mapstructure:"remote"`

	handle func(b []byte)
}

func (c *Config) Write(b []byte) (n int, err error) {
	if c.handle != nil {
		c.handle(b)
	}
	return len(b), nil
}

func (c *Config) WithLoggerHandle(handle func(b []byte)) {
	c.handle = handle
}

func New(cfg *Config) *zap.Logger {
	var cores []zapcore.Core
	if cfg.Console {
		cores = append(cores, internal.NewConsoleCore())
	}
	if cfg.Remote {
		cores = append(cores, internal.NewJsonCore(cfg))
	}
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	return logger
}
