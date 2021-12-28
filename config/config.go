package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("can't open config file:s", err)
	}
}

func GetConfig() *viper.Viper {
	return viper.GetViper()
}
