package utils

import "errors"

func ValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}

func ValidateRole(role string) error {
	validRoles := map[string]bool{
		"admin": true,
		"user":  true,
	}
	if !validRoles[role] {
		return errors.New("invalid role")
	}
	return nil
}
