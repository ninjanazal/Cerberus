package db_config

import "os"

// RedisConfigData represents the configuration data for a Redis connection.
type RedisConfigData struct {
	Address string
}

// DefaultRedisConfig is a global variable holding the default Redis configuration.
var DefaultRedisConfig RedisConfigData

func init() {
	DefaultRedisConfig.Address = "localhost:6379"
}

// ParseLineData parses a key-value pair and updates the corresponding field in the RedisConfigData struct.
// It matches the key against predefined fields and updates the matching field with the provided value.
//
// Parameters:
//   - p_key: A string representing the configuration key.
//   - p_value: A string representing the value to be set for the given key.
func (cfg *RedisConfigData) ParseLineData(p_key string, p_value string) {
	fMap := map[string]*string{
		"REDIS_ADDRESS": &cfg.Address,
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
