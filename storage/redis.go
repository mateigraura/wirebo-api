package storage

import (
	"github.com/mateigraura/wirebo-api/utils"
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func CreateRedisClient() {
	host := utils.GetEnvFile()[utils.RedisHost]
	opt, err := redis.ParseURL(host)
	if err != nil {
		log.Fatal(err)
	}

	Redis = redis.NewClient(opt)
}
