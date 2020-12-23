package apiserver

// Config type contains configuration info fields
type Config struct {
	BindAddr string `toml:"bind_addr"`
}

// NewConfig function creates new config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
