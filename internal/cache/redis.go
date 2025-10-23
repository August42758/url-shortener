package cache

import (
	"context"
	"urlShortener/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisConnect.Addres,
		Password: config.AppConfig.RedisConnect.Password,
		DB:       config.AppConfig.RedisConnect.Db,
	})

	if _, err := db.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return db, nil
}
