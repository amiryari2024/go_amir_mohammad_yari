package main

import (
	"fmt"
)

// User struct represents a user with an email and password
type User struct {
	Email    string
	Password string
}

// UserRepository struct represents a repository of users
type UserRepository struct {
	database []User
}

// Register adds a new user to the repository
func (r *UserRepository) Register(u User) error {
	if u.Email == "" || u.Password == "" {
		return fmt.Errorf("Register failed: Email or Password cannot be empty")
	}

	// Here you would hash the password before storing it
	r.database = append(r.database, u)
	fmt.Println("User registered successfully")
	return nil
}

// Login checks user credentials and returns an auth token if successful
func (r *UserRepository) Login(u User) (string, error) {
	if u.Email == "" || u.Password == "" {
		return "", fmt.Errorf("Login failed: Email or Password cannot be empty")
	}

	// Here you would hash the provided password and compare it with the stored hash
	for _, user := range r.database {
		if user.Email == u.Email && user.Password == u.Password {
			// Generate and return a unique auth token in a real application
			return "auth token", nil
		}
	}

	return "", fmt.Errorf("Login failed: Invalid email or password")
}

func main() {
	// Example usage
	repo := &UserRepository{}

	// Register a new user
	err := repo.Register(User{Email: "user@example.com", Password: "password123"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Attempt to log in
	token, err := repo.Login(User{Email: "user@example.com", Password: "password123"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Login successful, token:", token)
}
