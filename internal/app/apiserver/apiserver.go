package apiserver

import "github.com/sirupsen/logrus"

// APIServer type
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New func
func New(c *Config) *APIServer {
	return &APIServer{
		config: c,
		logger: logrus.New(),
	}
}

// Start func
func (s *APIServer) Start() error {
	err := s.loggerConf()
	if err != nil {
		return err
	}
	s.logger.Info("Starting api server..")
	return nil 
}

func (s *APIServer) loggerConf() error {
	lvl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(lvl)
	return nil
}
