package core

import (
	"context"
	"encoding/json"
	pb "github.com/lhdhtrc/logger-go/dep/server/v1"
	"github.com/lhdhtrc/logger-go/model"
	"time"
)

type LoggerCoreEntity struct {
	model.ConfigEntity
	sc pb.ServerLoggerServiceClient
}

func (s *LoggerCoreEntity) Write(b []byte) (n int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var row pb.AddRequest
	_ = json.Unmarshal(b, &row)
	row.AppId = s.AppId

	_, _ = s.sc.Add(ctx, &row)

	return len(b), nil
}
