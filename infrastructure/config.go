package infrastructure

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Database struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		Name        string `yaml:"db"`
		Connections int    `yaml:"connections"`
	} `yaml:"database"`
	RestServer struct {
		Enabled bool   `yaml:"enable"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
	} `yaml:"rest-server"`
	JsonRPRCServer struct {
		Enabled bool   `yaml:"enable"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
	} `yaml:"json-rpc-server"`
}

func GetConfig() Config {
	var config Config

	err := cleanenv.ReadConfig("config.yaml", &config)

	if err != nil {
		panic(err)
	}

	return config
}
