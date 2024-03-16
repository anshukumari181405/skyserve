// db/mysql.go
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "Akash@123"
	hostname = "localhost"
	port     = "3306"
	dbname   = "skyserve"
)

var db *sql.DB

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
}

// CreateUser adds a new user to the database
func (db *UserDatabase) CreateUser(user User) error {
	// Prepare statement for inserting data
	stmt, err := db.Prepare("INSERT INTO users(email, password) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail retrieves a user by email from the database
func (db *UserDatabase) GetUserByEmail(email string) (*User, error) {
	var user User
	// Query to select user by email
	err := db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
