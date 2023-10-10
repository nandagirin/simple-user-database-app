package database

import (
	"fmt"
	"sync"
	"user/models"
)

var (
	db []*models.User
	mu sync.Mutex
)

// Connect with database
func Connect() {
	db = make([]*models.User, 0)
	fmt.Println("Connected with Database")
}

// Insert user to database
func Insert(user *models.User) {
	mu.Lock()
	db = append(db, user)
	mu.Unlock()
}

// Get list of users from database
func Get() []*models.User {
	return db
}
