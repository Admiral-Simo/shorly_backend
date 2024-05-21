package db

import (
	"testing"

	"github.com/Admiral-Simo/shortly_backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUrlStore_GetUrls(t *testing.T) {
	// Set up a temporary SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Url{}); err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	// Create a new UrlStore instance
	store := NewUrlStore(db)

	// Create some test data
	userID := 1
	testUrls := []*models.Url{
		{ID: "hash1", URL: "https://example.com/1", UserID: userID},
		{ID: "hash2", URL: "https://example.com/2", UserID: userID},
	}

	// Insert test data into the database
	for _, u := range testUrls {
		if _, err := store.CreateUrl(u.UserID, u.URL, u.ID); err != nil {
			t.Fatalf("failed to create test url: %v", err)
		}
	}

	// Test GetUrls function
	urls, err := store.GetUrls(userID)
	if err != nil {
		t.Fatalf("failed to get urls: %v", err)
	}

	// Check if the retrieved URLs match the expected URLs
	if len(urls) != len(testUrls) {
		t.Fatalf("expected %d urls, got %d", len(testUrls), len(urls))
	}

	for i := range testUrls {
		if urls[i].ID != testUrls[i].ID || urls[i].URL != testUrls[i].URL || urls[i].UserID != testUrls[i].UserID {
			t.Fatalf("expected url %+v, got %+v", testUrls[i], urls[i])
		}
	}
}

func TestUrlStore_CreateUrl(t *testing.T) {
	// Set up a temporary SQLite database for testing
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.Url{}); err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	// Create a new UrlStore instance
	store := NewUrlStore(db)

	// Test CreateUrl function
	userID := 1
	url := "https://example.com/test"
	hash := "hash3"

	newUrl, err := store.CreateUrl(userID, url, hash)
	if err != nil {
		t.Fatalf("failed to create url: %v", err)
	}

	// Check if the created URL matches the expected URL
	if newUrl.ID != hash || newUrl.URL != url || newUrl.UserID != userID {
		t.Fatalf("expected url %+v, got %+v", &models.Url{ID: hash, URL: url, UserID: userID}, newUrl)
	}
}
