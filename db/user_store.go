package db

import (
	"errors"

	"github.com/Admiral-Simo/shortly_backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserStorer interface {
	CheckUser(username string, password string) (*models.User, error)
	CreateUser(username string, password string) (*models.User, error)
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{
		db: db,
	}
}

func (s *userStore) CheckUser(username string, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	valid := checkPassword(password, user.Password)

	if !valid {
		return nil, errors.New("invalid password")
	}

	user.Password = ""

	return &user, nil
}

func (s *userStore) CreateUser(username string, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err == nil {
		return nil, errors.New("username already taken")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user = models.User{
		Username: username,
		Password: hashedPassword,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	user.Password = ""

	return &user, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func checkPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
