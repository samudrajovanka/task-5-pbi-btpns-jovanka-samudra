package services

import (
	"photo-api/app"
	"photo-api/database"
	"photo-api/helpers"

	"golang.org/x/crypto/bcrypt"
)

func UserRegister(body *app.User) {
	if err := database.DB.Where("email=?", body.Email).First(&body).Error; err == nil {
		helpers.BadRequestError("User already registered")
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	body.Password = string(passwordHash)

	if err := database.DB.Create(&body).Error; err != nil {
		helpers.ServerErrorError(err.Error())
	}
}

func UserLogin(body app.Login) string {
	var user app.User

	if err := database.DB.Table("users").Where("email=?", body.Email).First(&user).Error; err != nil {
		helpers.UnauthorizedError("Email or password incorrect")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		helpers.UnauthorizedError("Email or password incorrect")
	}

	accessToken := helpers.GenerateAccessToken(user.Username)

	return accessToken
}

func UpdateByIdUser(id uint, body, user *app.User) {
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		helpers.NotFoundError(err.Error())
	}

	if user.ID != body.ID {
		helpers.ForbiddenError("You don't have a access")
	}

	if body.Password != "" {
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
		body.Password = string(passwordHash)
	}

	if err := database.DB.Model(&user).Where("id = ?", id).Updates(&body).Error; err != nil {
		helpers.ServerErrorError(err.Error())
	}
}

func DeleteByIdUser(id uint, user *app.User, userId uint) {
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		helpers.NotFoundError(err.Error())
	}

	if user.ID != userId {
		helpers.ForbiddenError("You don't have a access")
	}

	if err := database.DB.Delete(&user, id).Error; err != nil {
		helpers.ServerErrorError(err.Error())
	}
}
