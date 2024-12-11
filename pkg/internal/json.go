package internal

import (
	"github.com/lhdhtrc/logger-go/pkg/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func NewJsonCore(config core.LoggerConfig) zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format(time.DateTime))
	}
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "path"
	encoderConfig.TimeKey = "created_at"

	enc := zapcore.NewJSONEncoder(encoderConfig)
	writeSync := zapcore.AddSync(config)
	core := zapcore.NewCore(enc, writeSync, zap.NewAtomicLevel())

	return core
}
