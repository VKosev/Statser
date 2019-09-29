package main

import (
	"errors"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

//validateUser() returns error if validation of the string fails.
func validateUser(s string) (bool, error) {
	// Checking if the input is safe.
	if m, _ := regexp.MatchString("^[a-zA-Z]+[a-zA-Z0-9._-]*[a-zA-Z0-9]$", s); !m {
		err := errors.New("Invalid username")
		return false, err
	}
	//Checking if the input is atlest 5 characters long.
	if strings.Count(s, "") < 6 {
		err := errors.New("Username is less than 5 characters")
		return false, err

	}
	//Checking if such a username allready exists.
	_, err := db.getUsername(s)
	if err != nil {
		return true, nil
	}

	err = errors.New("Username allready exist")
	return false, err

}

func validatePsw(psw string) error {
	if strings.Count(psw, "") < 7 {
		err := errors.New("Password must be at least 6 characters")
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkLogin(user string, psw string) error {
	hash, err := db.getPassword(user)
	if err != nil {
		err := errors.New("Invalid password or username")
		return err
	}
	if checkPasswordHash(psw, hash) != true {
		err := errors.New("Invalid password or username")
		return err
	}
	return nil
}
