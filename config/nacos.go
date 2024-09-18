package config

type Nacos struct {
	IpAddr      string `json:"ipAddr" yaml:"ipAddr" mapstructure:"ipAddr"`
	Port        uint64 `json:"port" yaml:"port" mapstructure:"port"`
	ContextPath string `json:"contextPath" yaml:"contextPath" mapstructure:"contextPath"`
	Group       string `json:"group" yaml:"group" mapstructure:"group"`
	DataId      string `json:"dataId" yaml:"dataId" mapstructure:"dataId"`
	NamespaceId string `json:"namespaceId" yaml:"namespaceId" mapstructure:"namespaceId"`
	UserName    string `json:"userName" yaml:"userName" mapstructure:"userName"`
	Password    string `json:"password" yaml:"password" mapstructure:"password"`
}
