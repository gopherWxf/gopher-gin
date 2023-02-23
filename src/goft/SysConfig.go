package goft

import (
	"gopkg.in/yaml.v3"
	"log"
)

type ServerConfig struct {
	Port int32
	Name string
}
type SysConfig struct {
	Server *ServerConfig
}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{
		Port: 80,
		Name: "my web",
	}}
}
func InitConfig() *SysConfig {
	config := NewSysConfig()
	if b := LoadConfigFile(); b != nil {
		err := yaml.Unmarshal(b, config)
		if err != nil {
			log.Fatal(err)
		}
	}
	return config
}
