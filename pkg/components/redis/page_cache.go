package redis

import "github.com/redis/go-redis/v9"

type PageCache struct {
	Redis                   *redis.Client
	IsCacheEnabled          bool
	CacheVersion            string
	UserContentCacheVersion string
}

func (service *PageCache) IsEnabled() bool {
	return service.IsCacheEnabled
}
