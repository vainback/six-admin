package utils

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (password Password) TrimSpace() Password {
	return Password(strings.TrimSpace(string(password)))
}

func (password Password) Hash() Password {
	if password == "" {
		return password
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return Password(hash)
}

func (password Password) Verify(hash Password) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (password Password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}
