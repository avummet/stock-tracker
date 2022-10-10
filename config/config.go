package config

import (
	"errors"
	//"fmt"
	//"strings"
	"github.com/spf13/viper"
)

type Config struct {
	Ticker  string
	Numdays int
}

type ServerConfig struct {
	Config
	Port        int
	BindAddress string
}

func InitConfig(appName string) (Config, error) {
	var config Config

	//read env vars
	viper.SetEnvPrefix(appName)
	viper.AutomaticEnv()

	config.Ticker = viper.GetString("symbol")
	config.Numdays = viper.GetInt("ndays")

	if config.Numdays == 0 || config.Ticker == "" {
		return config, errors.New("ticker or port are missing")
	}

	return config, nil
}

func InitServerConfig(Config Config) (ServerConfig, error) {
	var servercfg ServerConfig

	servercfg.Config = Config
	servercfg.Port = viper.GetInt("port")
	servercfg.BindAddress = viper.GetString("bind_address")

	if servercfg.Port == 0 {
		return servercfg, errors.New("port needs to be provided")
	}

	if servercfg.BindAddress == "" {
		return servercfg, errors.New("BindAddress is empty")
	}

	return servercfg, nil
}
