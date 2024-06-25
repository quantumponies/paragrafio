package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReadConfig() {
	viper.SetDefault("environment", "development")
	viper.SetEnvPrefix("paragrafio")
	err := viper.BindEnv("environment")
	if err != nil {
		log.Fatalf("Error while binding env: %s", err)
	}
	viper.AutomaticEnv()
	viper.SetConfigName(viper.GetString("environment"))
	viper.AddConfigPath("/etc/paragrafio/configs")
	viper.AddConfigPath("./configs")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while loading config: %s", err)
	}
}