package cache

import (
	"github.com/go-redis/redis"
)

func NewRedisClient(host, port string) *redis.Client {
	// реализуйте создание клиента для Redis
	return redis.NewClient(
		&redis.Options{
			Addr: host + ":" + port,
		})
}
