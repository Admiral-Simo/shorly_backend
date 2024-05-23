package db

import (
	"github.com/Admiral-Simo/shortly_backend/models"
	"github.com/Admiral-Simo/shortly_backend/tools"
	"gorm.io/gorm"
)

type UrlStorer interface {
	GetUrls(userID int) ([]*models.Url, error)
	GetUrl(hash string) (*models.Url, error)
	CreateUrl(userID int, url string) (*models.Url, error)
}

type urlStore struct {
	db *gorm.DB
}

func NewUrlStore(db *gorm.DB) *urlStore {
	return &urlStore{db: db}
}

func (s *urlStore) GetUrls(userID int) ([]*models.Url, error) {
    var urls []*models.Url
    if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&urls).Error; err != nil {
        return nil, err
    }
    return urls, nil
}

func (s *urlStore) CreateUrl(userID int, url string) (*models.Url, error) {
	// make a random 6 digit hash here
	hash := tools.CreateUrlHash()
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

func (s *urlStore) GetUrl(hash string) (*models.Url, error) {
	var url *models.Url
	if err := s.db.Where("id = ?", hash).First(&url).Error; err != nil {
		return nil, err
	}
	return url, nil
}
