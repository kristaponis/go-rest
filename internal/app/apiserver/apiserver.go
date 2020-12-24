package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer type
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New func
func New(c *Config) *APIServer {
	return &APIServer{
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start func
func (s *APIServer) Start() error {
	err := s.loggerConf()
	if err != nil {
		return err
	}
	s.routerConf()
	s.logger.Info("Starting api server..")

	return http.ListenAndServe(s.config.BindAddr, s.router) 
}

// loggerConf sets parses log level and returns err if something goes wrong
func (s *APIServer) loggerConf() error {
	lvl, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(lvl)
	return nil
}

// routerConf contains router info
func (s *APIServer) routerConf() {
	s.router.HandleFunc("/hello", s.helloHandler())
}

func (s *APIServer) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
