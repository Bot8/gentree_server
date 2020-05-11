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
	JsonRPRCServer struct {
		Enabled bool   `yaml:"enable"`
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
	} `yaml:"json-rpc-server"`
	Encryption struct {
		JWTSecret string `yaml:"jwt-secret"`
	} `yaml:"encryption"`
}

func GetConfig() Config {
	var config Config

	err := cleanenv.ReadConfig("config.yaml", &config)

	if err != nil {
		panic(err)
	}

	return config
}
