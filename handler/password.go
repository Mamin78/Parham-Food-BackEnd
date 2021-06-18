package handler

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

func CheckPasswordLever(pass string) error {
	pass = strings.ToLower(pass)
	if len(pass) < 8 {
		return fmt.Errorf("password len is < 8")
	}
	num := `[0-9]{1}`
	aToz := `[a-z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, _ := regexp.MatchString(num, pass); !b {
		return fmt.Errorf("password need number")
	}
	if b, _ := regexp.MatchString(aToz, pass); !b {
		return fmt.Errorf("password need charachter")
	}
	if b, _ := regexp.MatchString(symbol, pass); b {
		return fmt.Errorf("password doesn't need symbol")
	}
	return nil
}

func PasswordToHash(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password can not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func PasswordsAreSame(original, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(original), []byte(plain))
	return err == nil
}
