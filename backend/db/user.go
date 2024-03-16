// db/user.go
package db

import (
	"errors"
	"regexp"
	"sync"
)

// User represents a user in the system
type User struct {
	ID       string
	Email    string
	Password string
}

// UserDatabase represents a database for managing users
type UserDatabase struct {
	users map[string]User // userID -> User
	mu    sync.RWMutex    // Mutex for concurrent access
}

// NewUserDatabase creates a new UserDatabase instance
func NewUserDatabase() *UserDatabase {
	return &UserDatabase{
		users: make(map[string]User),
	}
}

// CreateUser adds a new user to the database
func (db *UserDatabase) CreateUser(user User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	// Validate email format
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}

	// Validate password strength
	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Check if user with the same email already exists
	for _, existingUser := range db.users {
		if existingUser.Email == user.Email {
			return errors.New("user with this email already exists")
		}
	}

	// Generate unique ID for the user
	user.ID = generateID()

	// Save user to the database
	db.users[user.ID] = user
	return nil
}

// GetUserByEmail retrieves a user by email from the database
func (db *UserDatabase) GetUserByEmail(email string) (*User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// Find user by email
	for _, user := range db.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

// isValidEmail checks if the given email is in a valid format
func isValidEmail(email string) bool {
	// Regular expression for email validation
	// This is a basic validation, you may need to use a more comprehensive regex for production
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
