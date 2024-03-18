package config

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Redis    RedisConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type ServerConfig struct {
	AppName string
	Port    int
}

type RedisConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     27017,
			Username: "",
			Password: "",
		},
		Server: ServerConfig{
			AppName: "Ecommerce API",
			Port:    3000,
		},
		Redis: RedisConfig{
			Host:     "localhost",
			Port:     6379,
			Username: "",
			Password: "",
		},
	}
}
