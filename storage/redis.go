package storage

import (
	utils2 "github.com/mateigraura/wirebo-api/core/utils"
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func CreateRedisClient() {
	host := utils2.GetEnvFile()[utils2.RedisHost]
	opt, err := redis.ParseURL(host)
	if err != nil {
		log.Fatal(err)
	}

	Redis = redis.NewClient(opt)
}
