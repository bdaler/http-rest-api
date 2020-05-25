package store

// Config ...
type Config struct {
	DatabaseURL string `toml: "database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		DatabaseURL: "host=127.0.0.1 port=5431 user=onboarding password=onboarding dbname=onboarding sslmode=disable",
	}
}
