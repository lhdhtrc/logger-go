package core

import "fmt"

type LoggerOptions struct {
	Addr    string `json:"addr" yaml:"addr" mapstructure:"addr"`
	AppId   string `json:"app_id" yaml:"app_id" mapstructure:"app_id"`
	Console bool   `json:"console" yaml:"console" mapstructure:"console"`
	Remote  bool   `json:"remote" yaml:"remote" mapstructure:"remote"`
}

func (s *LoggerOptions) Write(b []byte) (n int, err error) {
	// TODO 远程存储调用
	fmt.Print(string(b), s.Addr, s.AppId)
	return len(b), nil
}
