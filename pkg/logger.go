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

func New(config *Config, handle func(b []byte)) *zap.Logger {
	config.handle = handle

	var cores []zapcore.Core
	if config.Console {
		cores = append(cores, internal.NewConsoleCore())
	}
	if config.Remote {
		cores = append(cores, internal.NewJsonCore(config))
	}
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	return logger
}
