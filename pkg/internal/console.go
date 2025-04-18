package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func NewConsoleCore() zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.ConsoleSeparator = " "
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "path"
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + t.Format(time.DateTime) + "]")
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
