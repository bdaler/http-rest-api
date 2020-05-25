package apiserver

// Config ...
type Config struct {
	BindAdrr string `toml: "bind_addr"`
	LogLevel string `toml: "log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAdrr: ":8080",
		LogLevel: "debug",
	}
}