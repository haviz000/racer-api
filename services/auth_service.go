package services

import (
	"errors"
	"log"
	"os"
)

func Login(username, password string) error {
	usernameEnv := os.Getenv("USERNAME_LOGIN_ENV")
	passwordEnv := os.Getenv("PASSWORD_LOGIN_ENV")
	log.Println("password", passwordEnv, "kocak", usernameEnv)

	if usernameEnv == "" || passwordEnv == "" {
		return errors.New("username or password is not set")
	}

	if username != usernameEnv || password != passwordEnv {
		return errors.New("invalid credentials")
	}

	return nil
}
