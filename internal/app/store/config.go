package store

// Config type
type Config struct {
	DatabaseURL string `toml:"database_url"`
}

// NewConfig function
func NewConfig() *Config {
	return &Config{}
}
