package db

import (
	"github.com/Admiral-Simo/shortly_backend/models"
	"gorm.io/gorm"
)

type UrlStorer interface {
	GetUrls(userID int) ([]*models.Url, error)
	CreateUrl(userID int, url string, hash string) (*models.Url, error)
}

type urlStore struct {
	db *gorm.DB
}

func NewUrlStore(db *gorm.DB) *urlStore {
	return &urlStore{db: db}
}

func (s *urlStore) GetUrls(userID int) ([]*models.Url, error) {
	var urls []*models.Url
	if err := s.db.Where("user_id = ?", userID).Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (s *urlStore) CreateUrl(userID int, url string, hash string) (*models.Url, error) {
	newUrl := &models.Url{
		ID:     hash,
		URL:    url,
		UserID: userID,
	}
	if err := s.db.Create(newUrl).Error; err != nil {
		return nil, err
	}
	return newUrl, nil
}
