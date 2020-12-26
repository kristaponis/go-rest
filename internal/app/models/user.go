package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"golang.org/x/crypto/bcrypt"
)

// User type
type User struct {
	ID          string
	Email       string
	Password    string
	EncPassword string
}

// Validate method
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u, 
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(validation.RuleFunc(requiredIf(u.EncPassword == ""))), validation.Length(6, 100)),
	)
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
