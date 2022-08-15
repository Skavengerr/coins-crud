package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"DBUSER"`
	Password string `mapstructure:"PASSWORD"`
	DbName   string `mapstructure:"DBNAME"`
}

func InitViper(path string) (cfg Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&cfg)
	return
}
