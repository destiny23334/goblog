package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"goblog/config"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Println(err)
		}
	}

	if err := viper.Unmarshal(&config.GlobalConfig); err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}
