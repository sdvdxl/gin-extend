package cache

import (
	"github.com/gin-gonic/contrib/cache"
	"github.com/sdvdxl/gin-extend/sources/util/config"
	"github.com/sdvdxl/gin-extend/sources/util/constant"
	. "github.com/sdvdxl/gin-extend/sources/util/log"
	"time"
)

var (
	Cache cache.CacheStore
)

func init() {
	Logger.Info("init cache ...")
	Cache = cache.NewRedisCache(config.AuthPageConfig.RedisHost, config.AuthPageConfig.RedisPassword, constant.LOGIN_EXPIRED_TIME)
	if err := Cache.Set("_test_is_valid_x_", "x", time.Second); err != nil {
		Logger.Error("error redis status,%v, will use memcached", err)
		Cache = cache.NewMemcachedStore(config.AuthPageConfig.MemcachedHosts, constant.LOGIN_EXPIRED_TIME)
		if err = Cache.Set("_test_is_valid_x_", "x", time.Second); err != nil {
			Logger.Error("error redis status,%v, will use memry cache", err)
			Cache = cache.NewInMemoryStore(constant.LOGIN_EXPIRED_TIME)
		}
	}

	Logger.Info("cache inited")
}
