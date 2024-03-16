// db/id_generator.go
package db

import "strconv"

var (
	userIDCounter  = 1 // Counter for generating unique user IDs
	shapeIDCounter = 1 // Counter for generating unique shape IDs
)

// generateUserID generates a unique ID for the user
func generateUserID() string {
	// Use the userIDCounter as the ID and then increment it for the next ID
	id := userIDCounter
	userIDCounter++
	return strconv.Itoa(id)
}

// generateShapeID generates a unique ID for the shape
func generateShapeID() string {
	// Use the shapeIDCounter as the ID and then increment it for the next ID
	id := shapeIDCounter
	shapeIDCounter++
	return strconv.Itoa(id)
}
