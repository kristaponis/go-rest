package store

import (
	"database/sql"

	// import postgres driver
	_ "github.com/lib/pq"
)

// Store type
type Store struct {
	config   *Config
	db       *sql.DB
	userRepo *UserRepository
}

// New function initializes new Store object
func New(c *Config) *Store {
	return &Store{
		config: c,
	}
}

// Open function
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

// Close function
func (s *Store) Close() {
	s.db.Close()
}

// User function
func (s *Store) User() *UserRepository {
	if s.userRepo != nil {
		return s.userRepo
	}
	s.userRepo = &UserRepository{
		store: s,
	}
	return s.userRepo
}
