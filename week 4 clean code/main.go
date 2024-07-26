package main

import (
	"errors"
	"fmt"
)

type User struct {
	Email    string
	Password string
}

type UserRepo struct {
	DB []User
}

// Register adds a new user to the repository.
// Returns an error if the registration fails due to invalid input.
func (r *UserRepo) Register(u User) error {
	if u.Email == "" || u.Password == "" {
		// Return an error instead of printing to console
		return errors.New("registration failed: email and password must not be empty")
	}

	// Append the user to the repository
	r.DB = append(r.DB, u)
	return nil
}

// Login checks if the user credentials are valid.
// Returns an authentication token if successful, or an empty string otherwise.
func (r *UserRepo) Login(u User) (string, error) {
	if u.Email == "" || u.Password == "" {
		// Return an error instead of printing to console
		return "", errors.New("login failed: email and password must not be empty")
	}

	for _, user := range r.DB {
		if user.Email == u.Email && user.Password == u.Password {
			return "auth token", nil
		}
	}

	return "", errors.New("login failed: invalid email or password")
}

func main() {
	repo := UserRepo{}

	user := User{Email: "example@example.com", Password: "password123"}

	// Example usage of Register method
	if err := repo.Register(user); err != nil {
		fmt.Println(err)
	}

	// Example usage of Login method
	token, err := repo.Login(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Login successful, token:", token)
	}
}
