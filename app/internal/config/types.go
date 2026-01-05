package config

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name string
	Env  string
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type FeatureFlags struct {
	Features map[string]bool `mapstructure:"features"`
	Limits   map[string]int  `mapstructure:"limits"`
}
