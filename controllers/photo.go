package controllers

import (
	"net/http"
	"photo-api/app"
	"photo-api/helpers"
	"photo-api/services"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	user, _ := ctx.Get("user")

	var photo app.Photo

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		helpers.BadRequestError(err.Error())
	}

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		helpers.BadRequestError(err.Error())
	}

	photo.UserID = user.(app.User).ID
	services.CreatePhoto(&photo)

	ctx.JSON(http.StatusCreated, app.ResponseData{
		Success: true,
		Message: "Successfully create photo",
		Data:    photo,
	})
}

func GetAllPhoto(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	photos := services.GetAllPhoto()

	ctx.JSON(http.StatusOK, app.ResponseData{
		Success: true,
		Message: "Successfully retrieved all photos",
		Data:    photos,
	})
}

func GetByIdPhoto(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	photoIdStr, _ := ctx.Params.Get("photoId")

	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		panic("Invalid photo id")
	}

	photo := services.GetByIdPhoto(uint(photoId))

	ctx.JSON(http.StatusOK, app.ResponseData{
		Success: true,
		Message: "Successfully retrieved all photos",
		Data:    photo,
	})
}

func UpdateByIdPhoto(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	user, _ := ctx.Get("user")

	var photo app.Photo

	photoIdStr, _ := ctx.Params.Get("photoId")

	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		panic("Invalid photo id")
	}

	var photoBody app.Photo

	if err := ctx.ShouldBindJSON(&photoBody); err != nil {
		helpers.BadRequestError(err.Error())
	}

	photoBody.UserID = user.(app.User).ID

	services.UpdateByIdPhoto(uint(photoId), &photoBody, &photo)

	ctx.JSON(http.StatusOK, app.ResponseData{
		Success: true,
		Message: "Successfully update photo",
		Data:    photo,
	})
}

func DeleteByIdPhoto(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	user, _ := ctx.Get("user")

	photoIdStr, _ := ctx.Params.Get("photoId")

	photoId, err := strconv.Atoi(photoIdStr)
	if err != nil {
		panic("Invalid photo id")
	}

	var photo app.Photo

	services.DeleteByIdPhoto(uint(photoId), &photo, user.(app.User).ID)

	ctx.JSON(http.StatusOK, app.BaseResponse{
		Success: true,
		Message: "Successfully delete photo",
	})
}
