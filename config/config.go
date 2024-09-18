package config

type Config struct {
	Server   Server   `json:"server" yaml:"server" mapstructure:"server"`
	Nacos    Nacos    `json:"nacos" yaml:"nacos" mapstructure:"nacos"`
	Mysql    Mysql    `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	DingTalk DingTalk `json:"dingTalk" yaml:"dingTalk" mapstructure:"dingTalk"`
}
