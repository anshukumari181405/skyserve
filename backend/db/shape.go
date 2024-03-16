// db/shape.go
package db

import (
	"sync"
)

// Shape represents a geospatial shape
type Shape struct {
	ID      string
	UserID  string
	GeoJSON string
}

// ShapeDatabase represents a database for managing shapes
type ShapeDatabase struct {
	shapes map[string]Shape // shapeID -> Shape
	mu     sync.RWMutex     // Mutex for concurrent access
}

// NewShapeDatabase creates a new ShapeDatabase instance
func NewShapeDatabase() *ShapeDatabase {
	return &ShapeDatabase{
		shapes: make(map[string]Shape),
	}
}

// CreateShape adds a new shape to the database
func (db *ShapeDatabase) CreateShape(shape Shape) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Generate unique ID for the shape
	shape.ID = generateID()

	// Save shape to the database
	db.shapes[shape.ID] = shape
}

// GetShapesByUserID retrieves shapes by user ID from the database
func (db *ShapeDatabase) GetShapesByUserID(userID string) []Shape {
	db.mu.RLock()
	defer db.mu.RUnlock()

	var userShapes []Shape

	// Find shapes by user ID
	for _, shape := range db.shapes {
		if shape.UserID == userID {
			userShapes = append(userShapes, shape)
		}
	}

	return userShapes
}
