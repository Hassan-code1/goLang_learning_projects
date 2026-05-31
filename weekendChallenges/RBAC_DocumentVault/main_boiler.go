package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "strings" // Remove "_" after TODO 4
	_ "time"    // Remove "_" after TODO 1

	_ "github.com/golang-jwt/jwt/v5" // Remove "_" after JWT TODOs
)

// Secret key used for signing JWTs
// IMPORTANT: This is just a dummy key for educational purposes. 
// Never hardcode real cryptographic keys in your source code in production!
var jwtKey = []byte("x8A9b2F4c1D7e3B6a0C5d8E2f9A1b4C3")

// User represents a database record
type User struct {
	Username string
	Password string
	Role     string
}

// Temporary in-memory database
var mockDB = map[string]User{
	"alice":   {Username: "alice", Password: "password123", Role: "admin"},
	"bob":     {Username: "bob", Password: "password123", Role: "employee"},
	"charlie": {Username: "charlie", Password: "password123", Role: "guest"},
}

// Login request payload
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


// LoginHandler validates credentials and issues JWTs
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user, exists := mockDB[creds.Username]

	if !exists || user.Password != creds.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// TODO 1:
	// Create jwt.MapClaims containing:
	// - "username"
	// - "role"
	// - "exp"
	//
	// The token should expire:
	// 1 hour from now
	//
	// Then create the JWT
	//
	// Use:
	// jwt.SigningMethodHS256

	
	// TODO 2:
	// Sign the token using:
	// jwtKey
	//
	// If signing fails:
	// return http.StatusInternalServerError


	// TODO 3:
	// Set:
	// Content-Type -> application/json
	//
	// Return:
	// {
	//   "token": "your_jwt_here"
	// }


}

// rbacMiddleware validates JWTs and enforces role boundaries
func rbacMiddleware(
	allowedRoles []string,
	next http.HandlerFunc,
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// TODO 4:
		// Read the Authorization header.
		//
		// Example:
		// Authorization: Bearer eyJhbGciOi...
		//
		// If the header is missing:
		// return http.StatusUnauthorized
		//
		// Remove the "Bearer " prefix using:
		// strings.TrimPrefix()


		// TODO 5:
		// Parse and validate the JWT using:
		// jwt.Parse(...)
		//
		// Inside the key function:
		// verify the signing method is:
		// *jwt.SigningMethodHMAC
		//
		// Return:
		// jwtKey
		//
		// If:
		// - err != nil
		// OR
		// - token.Valid == false
		//
		// Return:
		// http.StatusUnauthorized


		// TODO 6:
		// Extract:
		// token.Claims
		//
		// Cast it into:
		// jwt.MapClaims
		//
		// Then extract:
		// "role"
		//
		// Cast the role into:
		// string


		// TODO 7:
		// Loop through:
		// allowedRoles
		//
		// If the user's role matches:
		// call:
		// next(w, r)
		//
		// Then immediately return.


		// TODO 8:
		// If the user does not have permission:
		// return http.StatusForbidden


	}
}

func main() {
	fmt.Println("=== RBAC Document Vault ===")
	fmt.Println("Listening on port 8080...")

	mux := http.NewServeMux()

	// Public login endpoint
	mux.HandleFunc("/login", LoginHandler)

	// TODO 9:
	// Protect the routes below using:
	// rbacMiddleware(...)
	//
	// "/docs/public"
	// Allowed Roles:
	// - admin
	// - employee
	// - guest
	//
	// "/docs/internal"
	// Allowed Roles:
	// - admin
	// - employee
	//
	// "/docs/classified"
	// Allowed Roles:
	// - admin

	mux.HandleFunc("/docs/public", PublicDocHandler)         // Update this line
	mux.HandleFunc("/docs/internal", InternalDocHandler)     // Update this line
	mux.HandleFunc("/docs/classified", ClassifiedDocHandler) // Update this line

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}


// DO NOT MODIFY THESE HANDLERS

func PublicDocHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		"PUBLIC: Welcome to the company, here is the handbook.\n",
	))
}

func InternalDocHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		"INTERNAL: Q3 Financial Reports and Roadmap data.\n",
	))
}

func ClassifiedDocHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(
		"CLASSIFIED: Admin access only. Operation Midnight architecture.\n",
	))
}


