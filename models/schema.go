package models

import "gorm.io/gorm"

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex;not null" json:"-"`
	Password string `gorm:"not null" json:"password"`
	Urls     []Url  `gorm:"foreignKey:UserID" json:"-"`
}

type Url struct {
	ID     string `gorm:"primaryKey" json:"id"`
	URL    string `gorm:"not null" json:"url"`
	UserID int    `gorm:"not null" json:"userID"`
	User   User   `gorm:"foreignKey:UserID" json:"-"`
}

// AutoMigrate will migrate the database schema for the models
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Url{})
	if err != nil {
		return err
	}
	return nil
}
