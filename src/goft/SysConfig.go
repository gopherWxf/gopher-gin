package goft

import (
	"gopkg.in/yaml.v3"
	"log"
)

type UserConfig map[string]interface{}
type ServerConfig struct {
	Port int32
	Name string
	Html string
}
type SysConfig struct {
	Server *ServerConfig
	Config UserConfig
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

// GetConfigValue 递归读取用户配置文件
func GetConfigValue(m UserConfig, prefix []string, index int) interface{} {
	key := prefix[index]
	if v, ok := m[key]; ok {
		if index == len(prefix)-1 { //到了最后一个
			return v
		} else {
			index = index + 1
			if mv, ok := v.(UserConfig); ok { //值必须是UserConfig类型
				return GetConfigValue(mv, prefix, index)
			} else {
				return nil
			}
		}
	}
	return nil
}
