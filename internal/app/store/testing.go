package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore function
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tb ...string) {
		if len(tb) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tb, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
