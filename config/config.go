package config

type Config struct {
	Server   Server   `json:"server" yaml:"server" mapstructure:"server"`
	Mysql    Mysql    `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	DingTalk DingTalk `json:"ding-talk" yaml:"ding-talk" mapstructure:"ding-talk"`
}
