package redis

import (
	"github.com/KananHasanov747/gofiber_weather_app/pkg/utils"
	"github.com/gofiber/storage/redis/v3"
)

var Store *redis.Storage

var (
	REDIS_HOST     = utils.Show("REDIS_HOST", "127.0.0.1")
	REDIS_PORT     = utils.Show("REDIS_PORT", 6379)
	REDIS_USER     = utils.Show("REDIS_USER", "")
	REDIS_PASSWORD = utils.Show("REDIS_PASSWORD", "")
)

func Init() {
	Store = redis.New(redis.Config{
		Host:     REDIS_HOST,
		Port:     REDIS_PORT,
		Username: REDIS_USER,
		Password: REDIS_PASSWORD,
	})
}
