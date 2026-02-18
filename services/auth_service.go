package services

import (
	"errors"
	"log"
	"os"
)

func Login(username, password string) error {
	usernameEnv := os.Getenv("USERNAME_LOGIN_ENV")
	passwordEnv := os.Getenv("PASSWORD_LOGIN_ENV")
	log.Println("dari depan", username, "pass", password)
	log.Println("password", passwordEnv, "kocak", usernameEnv)

	if usernameEnv == "" || passwordEnv == "" {
		return errors.New("username or password is not set")
	}

	log.Println("is user name", username == usernameEnv, "is pass", password == passwordEnv)
	if username != usernameEnv || password != passwordEnv {
		return errors.New("invalid credentials")
	}

	return nil
}
