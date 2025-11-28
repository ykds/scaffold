package config

import (
	"os"
	"scaffold/pkg/logger"
	"scaffold/pkg/mongodb"
	"scaffold/pkg/redis"
	"scaffold/pkg/tdengine"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Debug        bool   `json:"debug" yaml:"debug"`
	Port         string `json:"port" yaml:"port"`
	ReadTimeout  int    `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout int    `json:"write_timeout" yaml:"write_timeout"`
}

type Config struct {
	Server Server          `json:"server" yaml:"server"`
	Logger logger.Config   `json:"logger" yaml:"logger"`
	Redis  redis.Config    `json:"redis" yaml:"redis"`
	Taos   tdengine.Config `json:"taos" yaml:"taos"`
	Mongo  mongodb.Config  `json:"mongo" yaml:"mongo"`
}

func InitConfig(file string) *Config {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(err)
	}
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	cfg := &Config{}
	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
