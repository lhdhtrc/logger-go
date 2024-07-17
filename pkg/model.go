package loger

type ConfigEntity struct {
	Console bool `json:"console" bson:"console" yaml:"console" mapstructure:"console"`
	Remote  bool `json:"remote" bson:"remote" yaml:"remote" mapstructure:"remote"`
}

type CoreEntity struct {
	ConfigEntity
	remoteHandle func(b []byte)
}
