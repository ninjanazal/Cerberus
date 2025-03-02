package db_config

import (
	"cerberus/internal/tools/logger"
	"os"
	"time"
)

// RedisConfigData represents the configuration data for a Redis connection.
type RedisConfigData struct {
	Address            string
	JWTDuration        string
	RefreshJWTDuration string
}

// DefaultRedisConfig is a global variable holding the default Redis configuration.
var DefaultRedisConfig RedisConfigData

func init() {
	DefaultRedisConfig.Address = "localhost:6379"
	DefaultRedisConfig.JWTDuration = "15m"
	DefaultRedisConfig.RefreshJWTDuration = "1h"
}

// ParseLineData parses a key-value pair and updates the corresponding field in the RedisConfigData struct.
// It matches the key against predefined fields and updates the matching field with the provided value.
//
// Parameters:
//   - p_key: A string representing the configuration key.
//   - p_value: A string representing the value to be set for the given key.
func (cfg *RedisConfigData) ParseLineData(p_key string, p_value string) {
	fMap := map[string]*string{
		"REDIS_ADDRESS":        &cfg.Address,
		"JWT_DURATION":         &cfg.JWTDuration,
		"JWT_REFRESH_DURATION": &cfg.RefreshJWTDuration,
	}

	if f, ok := fMap[p_key]; ok {
		*f = p_value
	}
}

// GetPassword retrieves the Redis password from the environment variable.
//
// This method checks the "REDIS_PASSWORD" environment variable and returns its value.
// If the environment variable is not set, it returns an empty string.
//
// Returns:
//   - string: The Redis password as set in the environment, or an empty string if not set.
func (cfg *RedisConfigData) GetPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}

// GetJWTDuration returns the duration specified in the RedisConfigData as a time.Duration.
// If the Duration field cannot be parsed as an integer, it logs an error and returns
// the default duration from DefaultRedisConfig.
//
// The duration is assumed to be in minutes and is converted to a time.Duration value.
//
// Returns:
//   - time.Duration: The parsed duration in minutes as a time.Duration value.
func (cfg *RedisConfigData) GetJWTDuration() time.Duration {
	i, err := time.ParseDuration(cfg.JWTDuration)
	if err != nil {
		logger.Log("Failed to get duration, returning default", logger.ERROR)
		i, _ := time.ParseDuration(DefaultRedisConfig.JWTDuration)
		return i
	}

	return i
}

// GetRefreshJWTDuration returns the refresh JWT duration specified in the RedisConfigData as a time.Duration.
// If the RefreshJWTDuration field cannot be parsed as a valid duration string, it logs an error and returns
// the default refresh JWT duration from DefaultRedisConfig.
//
// The duration is parsed using time.ParseDuration, which accepts strings like "300ms", "1.5h" or "2h45m".
//
// Returns:
//   - time.Duration: The parsed refresh JWT duration as a time.Duration value.
func (cfg *RedisConfigData) GetRefreshJWTDuration() time.Duration {
	i, err := time.ParseDuration(cfg.RefreshJWTDuration)
	if err != nil {
		logger.Log("Failed to get duration, return default", logger.ERROR)
		i, _ := time.ParseDuration(DefaultRedisConfig.RefreshJWTDuration)
		return i
	}

	return i
}
