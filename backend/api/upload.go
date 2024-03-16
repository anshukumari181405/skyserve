// api/upload.go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var data struct {
		UserID  string `json:"userId"`
		GeoJSON string `json:"geoJSON"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Validate data (e.g., check if UserID is valid)
	// Implement your validation logic here

	// Save data to database (replace this with actual database logic)
	// For this example, we're just printing the data
	fmt.Println("Uploaded data:", data)

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data uploaded successfully"})
}
