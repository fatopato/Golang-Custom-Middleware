package services

import (
	"errors"

	"github.com/fatopato/custom-middlewares/entity"
)

var users = map[string]entity.User{}

var errInvalidPassword = errors.New("invalid password")
var errInvalidUsername = errors.New("invalid username")
var errUserAlreadyExists = errors.New("user already exists")

func ValidateUser(user entity.User) error {
	if usr, ok := users[user.Username]; ok {
		if user.Password != usr.Password {
			return errInvalidPassword			
		}		
	} else {
		return errInvalidUsername
	}
	return nil
	
}

func RegisterUser(user entity.User) error {
	if _, ok := users[user.Username]; ok {
		return errUserAlreadyExists
	}
	users[user.Username] = user
	return nil
}