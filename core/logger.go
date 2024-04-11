package core

import (
	"context"
	"encoding/json"
	pb "github.com/lhdhtrc/logger-go/dep/server/v1"
	"time"
)

type LoggerOptions struct {
	Addr    string `json:"addr" yaml:"addr" mapstructure:"addr"`
	AppId   string `json:"app_id" yaml:"app_id" mapstructure:"app_id"`
	Console bool   `json:"console" yaml:"console" mapstructure:"console"`
	Remote  bool   `json:"remote" yaml:"remote" mapstructure:"remote"`
	sc      pb.ServerLoggerServiceClient
}

func (s *LoggerOptions) Write(b []byte) (n int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var row pb.AddRequest
	_ = json.Unmarshal(b, &row)
	row.AppId = s.AppId

	_, _ = s.sc.Add(ctx, &row)

	return len(b), nil
}
