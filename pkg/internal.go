package loger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func (s *CoreEntity) Write(b []byte) (n int, err error) {
	if s.remoteHandle != nil {
		s.remoteHandle(b)
	}
	return len(b), nil
}

func consoleCore() zapcore.Core {
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

func jsonCore(loggerCore *CoreEntity) zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "path"
	encoderConfig.TimeKey = "created_at"

	enc := zapcore.NewJSONEncoder(encoderConfig)
	writeSync := zapcore.AddSync(loggerCore)
	core := zapcore.NewCore(enc, writeSync, zap.NewAtomicLevel())

	return core
}
