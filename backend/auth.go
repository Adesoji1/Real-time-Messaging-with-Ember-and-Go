package main

import (
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
    var err error
    // Replace with your connection string
    connStr := "user=username dbname=mydb sslmode=disable password=mypassword"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    // Optionally, check the connection
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
}



// Define the authenticateUser function for user authentication
func authenticateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Find user by username in the in-memory database (replace with your database logic)
	var foundUser User
	for _, user := range usersDB {
		if user.Username == credentials.Username {
			foundUser = user
			break
		}
	}

	if foundUser.ID == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Compare the provided password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = foundUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString([]byte("your-secret-key")) // Replace with your secret key
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"token": tokenString}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func tokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Replace with your secret key
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}