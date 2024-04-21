package core

import (
	"github.com/lhdhtrc/logger-go/model"
)

type LoggerCoreEntity struct {
	model.ConfigEntity
	remoteHandle func(b []byte)
}

func (s *LoggerCoreEntity) Write(b []byte) (n int, err error) {
	if s.remoteHandle != nil {
		s.remoteHandle(b)
	}
	return len(b), nil
}
