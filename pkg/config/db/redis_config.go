package db_config

type RedisConfigData struct {
	Address  string
	Password string
}

var DefaultRedisConfig RedisConfigData

func init() {
	DefaultRedisConfig.Address = "localhost:6379"
	DefaultRedisConfig.Password = ""
}
