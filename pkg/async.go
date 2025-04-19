package logger

import (
	"fmt"
)

type AsyncLogger struct {
	queue  chan []byte
	handle func(b []byte)
}

func NewAsyncLogger(size int, handle func(b []byte)) *AsyncLogger {
	logger := &AsyncLogger{
		queue:  make(chan []byte, size),
		handle: handle,
	}

	go logger.init()

	return logger
}

func (l *AsyncLogger) init() {
	for b := range l.queue {
		l.handle(b)
	}
}

func (l *AsyncLogger) Logger(b []byte) {
	select {
	case l.queue <- b:
	default:
		fmt.Println("Log queue is full, dropping log")
	}
}
