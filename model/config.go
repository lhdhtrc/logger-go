package model

type ConfigEntity struct {
	Console bool `json:"console" bson:"console" yaml:"console" mapstructure:"console"`
	Remote  bool `json:"remote" bson:"remote" yaml:"remote" mapstructure:"remote"`
}
