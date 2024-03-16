// api/signup.go
package api

import (
	"encoding/json"
	"net/http"

	"github.com/anshukumari181405/skyserve/db"
)

var userDB *db.UserDatabase

func init() {
	userDB = db.NewUserDatabase()
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var newUser db.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate user input
	if newUser.Email == "" || newUser.Password == "" {
		http.Error(w, "Email and password are required fields", http.StatusBadRequest)
		return
	}

	// Create user in the database
	err = userDB.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
