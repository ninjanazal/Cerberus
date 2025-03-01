package database

import (
	logger "cerberus/internal/tools/logger"
	"cerberus/pkg/config"
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

// RedisPack encapsulates a Redis client and its associated context.
type RedisPack struct {
	Client *redis.Client
	Ctx    context.Context
}

// ConnectRedis establishes a connection to Redis using the provided configuration.
// It creates a new RedisPack with a Redis client and a background context.
//
// Parameters:
//   - p_cfg: A pointer to the ConfigData structure containing Redis connection details.
//
// Returns:
//   - *RedisPack: A pointer to the created RedisPack if the connection is successful.
//   - error: An error if the connection fails, nil otherwise.
//
// The function attempts to ping the Redis server to verify the connection.
// It logs the connection status and returns an error if the connection fails.
func ConnectRedis(p_cfg *config.ConfigData) (*RedisPack, error) {
	var pack RedisPack = RedisPack{
		Ctx: context.Background(),
	}
	pack.Client = redis.NewClient(&redis.Options{
		Addr:     p_cfg.RedisData.Address,
		Password: p_cfg.RedisData.GetPassword(),
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
