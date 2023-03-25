package utils

import (
	"github.com/spf13/viper"
	"log"
)

const (
	configPath = "./config"
	configName = "config"
	configType = "yaml"
)

func GetConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath(configPath)
	config.SetConfigName(configName)
	config.SetConfigType(configType)

	if err := config.ReadInConfig(); err != nil {
		log.Panic("config err: " + err.Error())
	}
	return config
}
