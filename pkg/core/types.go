package core

// LoggerConfig 定义基础配置接口
type LoggerConfig interface {
	Write(p []byte) (n int, err error)
}
