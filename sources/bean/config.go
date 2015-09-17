package bean

import "github.com/sdvdxl/go-tools/collections"

type AuthPageConfig struct {
	LoginPages *collections.Set
	LoginPaths *collections.Set

	RedisHost           string
	RedisPassword       string
	RedisExpiredSeconds int
	MemcachedHosts      []string
}

type AuthPageConfigError struct {
}

func (this AuthPageConfig) Error() string {
	return "auth page config error"
}
