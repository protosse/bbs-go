package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	LogLevel string `yaml:"logLevel"`
	LogFile  string `yaml:"logFile"`
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	MaxSize  int64  `yaml:"maxSize"`
	Pprof    bool   `yaml:"pprof"`
	Cache    struct {
		Driver string `yaml:"driver"`
	} `yaml:"cache"`
	Limit struct {
		Disable bool    `yaml:"disable"`
		Limit   float64 `yaml:"limit"`
		Burst   int     `yaml:"burst"`
	} `yaml:"limit"`
	DB struct {
		Adapter string `yaml:"adapter"`
		Prefix  string `yaml:"prefix"`
		Conn    string `yaml:"conn"`
	} `yaml:"db"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

func Init(fileName string) (*Config, error) {
	config := &Config{}
	yamlFile, err := ioutil.ReadFile(fileName)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, config)
	}
	return config, err
}
