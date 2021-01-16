package tools

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	EtcdConfig   EtcdConfig   `yaml:"etcd"`
	JaegerConfig JaegerConfig `yaml:"jaeger"`
	AppConfig    AppConfig    `yaml:"app"`
	UserMicroName UserMicroName `yaml:"userMicroName"`
}

type EtcdConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type JaegerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}
type AppConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
	Name string `yaml:"name"`
}
type UserMicroName struct {
	Name string	`yaml:"name"`
}

var Cfg *Config

//GetConfig 主动获取配置
func GetConfig() *Config {
	return Cfg
}

//ParseConfig 解析config
func ParseConfig(path string) (cfg *Config) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("open file err,", err)
	}
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatalln("unmarshal fail err,", err)
	}
	Cfg = cfg
	return
}
