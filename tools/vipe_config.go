package tools

import (
	"CodeHive/logger"
	"github.com/spf13/viper"
)

func ConfigSet() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("tools")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Error(err)
		return Config{}, err
	}

	var configs Config
	if err := viper.Unmarshal(&configs); err != nil {
		logger.Log.Error(err)
		return Config{}, err
	}

	return configs, nil
}
