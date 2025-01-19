package models

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port      int
	JwtSecret string
	Salt      string
}
