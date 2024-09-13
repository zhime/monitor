package config

type DingTalk struct {
	Webhook string `json:"web-hook" yaml:"web-hook" mapstructure:"web-hook"`
	Secret  string `json:"secret" yaml:"secret" mapstructure:"secret"`
}
