package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func getConnectionString(configKey string) string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString(configKey+".host"),
		viper.GetString(configKey+".port"),
		viper.GetString(configKey+".name"),
		viper.GetString(configKey+".user"),
		viper.GetString(configKey+".password"),
	)
}

func Connect() *DB {
	connection, err := gorm.Open(postgres.Open(getConnectionString("db")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DB{Connection: connection}
}