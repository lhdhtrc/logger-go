package loger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(config *ConfigEntity, handle func(b []byte)) *zap.Logger {
	core := &CoreEntity{
		ConfigEntity: *config,
		handle:       handle,
	}

	var cores []zapcore.Core
	if core.Console {
		cores = append(cores, consoleCore())
	}
	if core.Remote {
		cores = append(cores, jsonCore(core))
	}
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	return logger
}
