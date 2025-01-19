package models

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port          int
	ReadTimeout   int
	WriteTimeout  int
	AuthPort      int
	ResourcesPort int
}
