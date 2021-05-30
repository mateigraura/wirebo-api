package storage

import (
	"github.com/go-redis/redis/v8"
	"github.com/mateigraura/wirebo-api/core/utils"
)

var Redis *redis.Client

func CreateRedisClient() {
	host := utils.GetEnvFile()[utils.RedisHost]
	opt, err := redis.ParseURL(host)
	if err != nil {
		panic(err)
	}

	Redis = redis.NewClient(opt)
}
