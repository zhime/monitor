package config

type Config struct {
	Mysql    Mysql    `json:"mysql" yaml:"mysql" mapstructure:"mysql"`
	DingTalk DingTalk `json:"ding-talk" yaml:"ding-talk" mapstructure:"ding-talk"`
}
