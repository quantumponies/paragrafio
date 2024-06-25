package main

import (
	"fmt"
	"net/http"
	"paragrafio/pkg/components/db"
	"paragrafio/pkg/components/redis"
	"paragrafio/pkg/config"
	"paragrafio/pkg/routers"
	"paragrafio/pkg/services"

	"github.com/spf13/viper"
)

func main() {
	config.ReadConfig()

	dbConnection := db.Connect()
	redisConnection := redis.Connect()

	transactionProvider := &db.TransactionProvider{
		DBConn:          dbConn.Connection,
		TimescaleDBConn: dbConn.TimescaleConnection,
	}

	providers := services.Providers{
		PageCache: &redis.PageCache{
			Redis:                   redisConnection,
			IsCacheEnabled:          viper.GetBool("redis.is_cache_enabled"),
			CacheVersion:            viper.GetString("redis.cache_version"),
			UserContentCacheVersion: viper.GetString("redis.user_content_cache_version"),
		},
	}
	port := fmt.Sprintf(":%s", viper.GetString("web.port"))
	http.ListenAndServe(port, routers.CreateBackendRouter(&providers))
}