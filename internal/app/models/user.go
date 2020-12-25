package models

import "golang.org/x/crypto/bcrypt"

// User type
type User struct {
	ID          string
	Email       string
	Password    string
	EncPassword string
}

// BeforeCreate method
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encString(u.Password)
		if err != nil {
			return err
		}
		u.EncPassword = enc
	}
	return nil
}

func encString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
