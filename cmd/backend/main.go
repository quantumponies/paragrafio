package main

import (
	"fmt"
	"net/http"
	"paragrafio/pkg/components/db"
	"paragrafio/pkg/components/redis"
	"paragrafio/pkg/components/s3"
	"paragrafio/pkg/config"
	"paragrafio/pkg/routers"
	"paragrafio/pkg/services"

	"github.com/spf13/viper"
)

func main() {
	config.ReadConfig()

	dbConnection := db.Connect()
	redisConnection := redis.Connect()
	s3Connection := s3.Init()

	transactionProvider := &db.TransactionProvider{
		DBConn: dbConnection.Connection,
	}

	providers := services.Providers{
		TransactionProvider: transactionProvider,
		PageCache: &redis.PageCache{
			Redis:                   redisConnection,
			IsCacheEnabled:          viper.GetBool("redis.is_cache_enabled"),
			CacheVersion:            viper.GetString("redis.cache_version"),
			UserContentCacheVersion: viper.GetString("redis.user_content_cache_version"),
		},
		S3Storage: &s3.S3Storage{S3: s3Connection},
	}
	port := fmt.Sprintf(":%s", viper.GetString("web.port"))
	http.ListenAndServe(port, routers.CreateBackendRouter(&providers))
}