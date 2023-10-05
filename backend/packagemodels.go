package main

// User struct represents user data
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// In-memory database
var usersDB []User
