package logger

import (
	"encoding/json"
	"github.com/lhdhtrc/logger-go/pkg/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMap = map[string]int{
	"info":  1,
	"warn":  3,
	"error": 4,
	"panic": 5,
	"debug": 6,
}

type log struct {
	Level     string `json:"level"`
	CreatedAt string `json:"created_at"`
	Path      string `json:"path"`
	Message   string `json:"message"`
	TraceId   string `json:"trace_id"`
}

type Conf struct {
	Console bool `json:"console" bson:"console" yaml:"console" mapstructure:"console"`
	Remote  bool `json:"remote" bson:"remote" yaml:"remote" mapstructure:"remote"`

	handle func(b []byte)
}

func (c *Conf) Write(b []byte) (n int, err error) {
	if c.handle != nil {

		var data log
		_ = json.Unmarshal(b, &data)

		l := make(map[string]interface{})

		l["Path"] = data.Path
		l["Level"] = levelMap[data.Level]
		l["Content"] = data.Message
		l["TraceId"] = data.TraceId
		l["CreatedAt"] = data.CreatedAt

		b, _ = json.Marshal(&l)

		c.handle(b)
	}
	return len(b), nil
}

func New(conf *Conf, handle func(b []byte)) *zap.Logger {
	conf.handle = handle

	var cores []zapcore.Core
	if conf.Console {
		cores = append(cores, internal.NewConsoleCore())
	}
	if conf.Remote {
		cores = append(cores, internal.NewJsonCore(conf))
	}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}
