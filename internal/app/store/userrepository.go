package store

import "github.com/kristaponis/go-rest/internal/app/models"

// UserRepository type
type UserRepository struct {
	store *Store
}

// Create function
func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := ur.store.db.QueryRow(
		"INSERT INTO users (id, email, enc_password) VALUES ($1, $2, $3) RETURNING id", u.ID, u.Email, u.EncPassword).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail function
func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	usr := &models.User{}
	if err := ur.store.db.QueryRow(
	"SELECT id, email, enc_password FROM users WHERE email = $1", email).Scan(&usr.ID, &usr.Email, &usr.EncPassword); err != nil {
		return nil, err
	}
	return usr, nil
}
