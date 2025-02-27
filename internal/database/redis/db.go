package db_redis

import (
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisPack struct {
	Client *redis.Client
	Ctx    context.Context
}

func ConnectRedis(p_cfg *config.ConfigData) (*RedisPack, error) {
	var pack RedisPack = RedisPack{
		Ctx: context.Background(),
	}
	pack.Client = redis.NewClient(&redis.Options{
		Addr:     p_cfg.RedisData.Address,
		Password: p_cfg.RedisData.Password,
		DB:       0,
	})

	_, err := pack.Client.Ping(pack.Ctx).Result()
	if err != nil {
		msg := "Failed to connect to Redis - " + err.Error()
		logger.Log(msg, logger.ERROR)

		return nil, errors.New(msg)
	}

	logger.Log("🧲 Connected to redis successfully", logger.INFO)
	return &pack, nil

}
