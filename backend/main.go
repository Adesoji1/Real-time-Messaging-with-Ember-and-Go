package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"backend/proto"
	"strconv"
	"sync"
    "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"time"
	"context"
	"github.com/hashicorp/terraform-exec/tfexec"
    //"github.com/rs/cors"
)



var mu sync.Mutex



func registerUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    mu.Lock()
    defer mu.Unlock()

    // Check if the username already exists
    for _, user := range usersDB {
        if user.Username == newUser.Username {
            http.Error(w, "Username already exists", http.StatusConflict)
            return
        }
    }

    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Store user in the in-memory database
    newUser.ID = "user_" + strconv.Itoa(len(usersDB)+1)
    newUser.Password = string(hashedPassword)
    usersDB = append(usersDB, newUser)

    // Avoid returning the password, even if it's hashed
    newUser.Password = ""
    if err := json.NewEncoder(w).Encode(newUser); err != nil {
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}


//lint:ignore U1000 reason: "This function will be used in future implementations"
func generateJWT(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    // Replace 'your_secret_key' with your actual secret key
    return token.SignedString([]byte("your_secret_key"))
}





// func main() {
// 	// Start the gRPC server in a separate goroutine

    
// 	go StartGRPCServer()

// 	setupDatabase()

// 	r := mux.NewRouter()

// 	r.HandleFunc("/register", registerUser).Methods("POST")
// 	r.HandleFunc("/login", authenticateUser).Methods("POST")
// 	// Protected route using a subrouter
// 	protectedRoute := r.PathPrefix("/protected").Subrouter()
// 	protectedRoute.Use(tokenValidationMiddleware)
// 	protectedRoute.HandleFunc("", protectedHandler).Methods("GET")

// 	http.Handle("/", r)

// 	log.Println("HTTP server started on :8080")
// 	http.ListenAndServe(":8080", nil)

// 	workingDir := "backend/terraform-configs"
//     execPath := "/usr/bin/terraform" // usually just "terraform" if it's in PATH

//     // Initialize Terraform
//     tf, err := tfexec.NewTerraform(workingDir, execPath)
//     if err != nil {
//         log.Fatalf("error running NewTerraform: %s", err)
//     }

//     // Apply Terraform configurations
//     err = tf.Apply(context.Background())
//     if err != nil {
//         log.Fatalf("error applying Terraform configurations: %s", err)
//     }

//     // ... (rest of your Go code)

// }

// ... [Your imports and other functions]

func main() {
	// Start the gRPC server in a separate goroutine
	go StartGRPCServer()

	// Set up the database
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("failed to setup database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/register", registerUser).Methods("POST")
	r.HandleFunc("/login", authenticateUser).Methods("POST")
	// Protected route using a subrouter
	protectedRoute := r.PathPrefix("/protected").Subrouter()
	protectedRoute.Use(tokenValidationMiddleware)
	protectedRoute.HandleFunc("", protectedHandler).Methods("GET")

	http.Handle("/", r)

	log.Println("HTTP server started on :8080")
	http.ListenAndServe(":8080", nil)

	workingDir := "backend/terraform-configs"
    execPath := "/usr/bin/terraform" // usually just "terraform" if it's in PATH

    // Initialize Terraform
    tf, err := tfexec.NewTerraform(workingDir, execPath)
    if err != nil {
        log.Fatalf("error running NewTerraform: %s", err)
    }

    // Apply Terraform configurations
    err = tf.Apply(context.Background())
    if err != nil {
        log.Fatalf("error applying Terraform configurations: %s", err)
    }

    // ... (rest of your Go code)

}





func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// This handler is protected by the tokenValidationMiddleware
	// Only requests with a valid JWT token can access this
	w.Write([]byte("Protected data"))
}

func StartGRPCServer() {
	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Register the chat service with the server
	messaging.RegisterChatServiceServer(grpcServer, &ChatServiceServer{})

	// Start listening for incoming gRPC connections
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}







