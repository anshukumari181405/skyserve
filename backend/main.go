// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// User struct represents a user in the system
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Shape struct represents a geospatial shape
type Shape struct {
	ID      string `json:"id"`
	UserID  string `json:"userId"`
	GeoJSON string `json:"geoJSON"`
}

var users []User
var shapes []Shape

func main() {
	http.HandleFunc("/api/signup", SignupHandler)
	http.HandleFunc("/api/login", LoginHandler)
	http.HandleFunc("/api/upload", UploadHandler)
	http.HandleFunc("/api/shapes", ShapesHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	body, err := io.ReadAll(io.Reader(r.Body))
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse JSON
	var newUser User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the user
	newUser.ID = generateID()

	// Save user to memory (in a real app, this would save to a database)
	users = append(users, newUser)

	// Return the newly created user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	body, err := io.ReadAll(io.Reader(r.Body))
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse JSON
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err = json.Unmarshal(body, &creds)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Find user by email and password (in a real app, this would check a database)
	var loggedInUser User
	for _, user := range users {
		if user.Email == creds.Email && user.Password == creds.Password {
			loggedInUser = user
			break
		}
	}

	// If user not found or password is incorrect
	if loggedInUser.ID == "" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return the logged-in user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loggedInUser)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	body, err := io.ReadAll(io.Reader(r.Body))
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse JSON
	var newShape Shape
	err = json.Unmarshal(body, &newShape)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the shape
	newShape.ID = generateID()

	// Save shape to memory (in a real app, this would save to a database)
	shapes = append(shapes, newShape)

	// Return the newly created shape
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newShape)
}

func ShapesHandler(w http.ResponseWriter, r *http.Request) {
	// Return all shapes in memory
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shapes)
}

func generateID() string {
	// Generate a random 8-character alphanumeric string for IDs
	return "id123456"
}
