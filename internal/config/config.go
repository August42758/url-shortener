package config

import (
	"github.com/spf13/viper"
)

const IsDebug bool = false

type Config struct {
	DbConnect struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Addres   string `mapstructure:"addres"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"db_connect"`
	ServerAddres string `mapstructure:"server_addres"`
}

var AppConfig Config

func LoadConfig() {
	if IsDebug {
		viper.AddConfigPath("../../internal/config")
	} else {
		viper.AddConfigPath("./internal/config")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(err)
	}
}
