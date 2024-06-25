package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	return client
}
