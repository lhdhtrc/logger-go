package model

type ConfigEntity struct {
	Addr    string `json:"addr" bson:"addr" yaml:"addr" mapstructure:"addr"`
	AppId   string `json:"app_id" bson:"app_id" yaml:"app_id" mapstructure:"app_id"`
	Console bool   `json:"console" bson:"console" yaml:"console" mapstructure:"console"`
	Remote  bool   `json:"remote" bson:"remote" yaml:"remote" mapstructure:"remote"`
}
