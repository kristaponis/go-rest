package apiserver

// APIServer type
type APIServer struct {
	config *Config
}

// New func
func New(c *Config) *APIServer {
	return &APIServer{
		config: c,
	}
}

// Start func
func (s *APIServer) Start() error {
	return nil
}
