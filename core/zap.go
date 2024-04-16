package core

import (
	"fmt"
	pb "github.com/lhdhtrc/logger-go/dep/server/v1"
	"github.com/lhdhtrc/logger-go/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"time"
)

func NewConsoleCore() zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "path"
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + t.Format("2006-01-02 15:04:05") + "]")
	}
	encoderConfig.EncodeLevel = func(t zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + t.String() + "]")
	}
	encoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + caller.TrimmedPath() + "]")
	}

	enc := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zap.NewAtomicLevel())

	return core
}

func NewJsonCore(loggerCore *LoggerCoreEntity) zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "path"

	enc := zapcore.NewJSONEncoder(encoderConfig)
	writeSync := zapcore.AddSync(loggerCore)
	core := zapcore.NewCore(enc, writeSync, zap.NewAtomicLevel())

	return core
}

func Setup(config model.ConfigEntity) *zap.Logger {
	core := &LoggerCoreEntity{ConfigEntity: config}

	var cores []zapcore.Core
	if core.Console {
		cores = append(cores, NewConsoleCore())
	}
	if core.Remote {
		if cli, err := grpc.Dial(core.Addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err == nil {
			core.sc = pb.NewServerLoggerServiceClient(cli)
			cores = append(cores, NewJsonCore(core))
		} else {
			fmt.Println("the grpc service cannot be connected")
		}
	}
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	return logger
}
