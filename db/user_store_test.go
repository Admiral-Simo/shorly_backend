package db

import (
	"testing"

	"github.com/Admiral-Simo/shortly_backend/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	store := NewUserStore(db)

	user, err := store.CreateUser("testuser", "testpassword")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)

	// Trying to create the same user again should result in an error
	_, err = store.CreateUser("testuser", "testpassword")
	assert.Error(t, err)
	assert.Equal(t, "username already taken", err.Error())
}

func TestCheckUser(t *testing.T) {
	db := setupTestDB(t)
	store := NewUserStore(db)

	_, err := store.CreateUser("testuser", "testpassword")
	assert.NoError(t, err)

	// Valid user and password
	user, err := store.CheckUser("testuser", "testpassword")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)

	// Invalid password
	_, err = store.CheckUser("testuser", "wrongpassword")
	assert.Error(t, err)
	assert.Equal(t, "invalid password", err.Error())

	// Non-existing user
	_, err = store.CheckUser("nonexistent", "testpassword")
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}
