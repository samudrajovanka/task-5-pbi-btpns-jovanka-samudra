package database

import "photo-api/models"

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
