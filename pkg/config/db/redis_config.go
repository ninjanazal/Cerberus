package db_config

// RedisConfigData represents the configuration data for a Redis connection.
type RedisConfigData struct {
	Address  string
	Password string
}

// DefaultRedisConfig is a global variable holding the default Redis configuration.
var DefaultRedisConfig RedisConfigData

func init() {
	DefaultRedisConfig.Address = "localhost:6379"
	DefaultRedisConfig.Password = ""
}

// ParseLineData parses a key-value pair and updates the corresponding field in the RedisConfigData struct.
// It matches the key against predefined fields and updates the matching field with the provided value.
//
// Parameters:
//   - p_key: A string representing the configuration key.
//   - p_value: A string representing the value to be set for the given key.
func (cfg *RedisConfigData) ParseLineData(p_key string, p_value string) {
	fMap := map[string]*string{
		"REDIS_ADDRESS":  &cfg.Address,
		"REDIS_PASSWORD": &cfg.Password,
	}

	if f, ok := fMap[p_key]; ok {
		*f = p_value
	}
}
