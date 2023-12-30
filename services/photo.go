package services

import (
	"photo-api/app"
	"photo-api/database"
	"photo-api/helpers"
)

func GetAllPhoto() (photos []app.Photo) {

	if err := database.DB.Find(&photos).Error; err != nil {
		helpers.ServerErrorError(err.Error())
	}

	return photos
}

func CreatePhoto(body *app.Photo) {
	if err := database.DB.Create(body).Error; err != nil {
		helpers.ServerErrorError(err.Error())
	}
}

func GetByIdPhoto(id uint) (photo app.Photo) {
	if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		helpers.NotFoundError("Photo not found")
	}

	return photo
}

func UpdateByIdPhoto(id uint, body, photo *app.Photo) {
	if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		helpers.NotFoundError(err.Error())
	}

	if photo.UserID != body.UserID {
		helpers.ForbiddenError("You don't have a access")
	}

	if err := database.DB.Model(&photo).Where("id = ?", id).Updates(&body).Error; err != nil {
		helpers.ServerErrorError(err.Error())
	}
}

func DeleteByIdPhoto(id uint, photo *app.Photo, userId uint) {
	if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		helpers.NotFoundError(err.Error())
	}

	if photo.UserID != userId {
		helpers.ForbiddenError("You don't have a access")
	}

	if err := database.DB.Delete(&photo, id).Error; err != nil {
		helpers.ServerErrorError(err.Error())
	}
}
