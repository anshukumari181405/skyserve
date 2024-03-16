// api/shape.go
package api

import (
	"encoding/json"
	"net/http"

	"github.com/anshukumari181405/skyserve/db"
)

var shapeDB *db.ShapeDatabase

func init() {
	shapeDB = db.NewShapeDatabase()
}

func ShapeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateShapeHandler(w, r)
	case http.MethodGet:
		GetShapesHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateShapeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var newShape db.Shape
	err := json.NewDecoder(r.Body).Decode(&newShape)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Create shape in the database
	shapeDB.CreateShape(newShape)

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Shape created successfully"})
}

func GetShapesHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from query parameter or request body
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve shapes by user ID from the database
	shapes := shapeDB.GetShapesByUserID(userID)

	// Return shapes as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shapes)
}
