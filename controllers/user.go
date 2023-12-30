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

func UserRegister(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	var user app.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.BadRequestError(err.Error())
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		helpers.BadRequestError(err.Error())
	}

	services.UserRegister(&user)

	ctx.JSON(http.StatusCreated, app.ResponseData{
		Success: true,
		Message: "Registration successful",
		Data: app.CreatedUserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	})
}

func UserLogin(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	var loginBody app.Login

	if err := ctx.ShouldBindJSON(&loginBody); err != nil {
		helpers.BadRequestError(err.Error())
	}

	if _, err := govalidator.ValidateStruct(loginBody); err != nil {
		helpers.BadRequestError(err.Error())
	}

	accessToken := services.UserLogin(loginBody)

	ctx.JSON(http.StatusOK, app.ResponseData{
		Success: true,
		Message: "Login successfully",
		Data: app.LoginResponse{
			AccessToken: accessToken,
		},
	})
}

func UpdateByIdUser(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	userReq, _ := ctx.Get("user")

	var user app.User

	userIdStr, _ := ctx.Params.Get("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		panic("Invalid user id")
	}

	var userBody app.User

	if err := ctx.ShouldBindJSON(&userBody); err != nil {
		helpers.BadRequestError(err.Error())
	}

	userBody.ID = userReq.(app.User).ID

	services.UpdateByIdUser(uint(userId), &userBody, &user)

	ctx.JSON(http.StatusOK, app.ResponseData{
		Success: true,
		Message: "Successfully update user",
		Data:    user,
	})
}

func DeleteByIdUser(ctx *gin.Context) {
	defer helpers.RecoverErrorResponse(ctx)

	userReq, _ := ctx.Get("user")

	userIdStr, _ := ctx.Params.Get("userId")

	photoId, err := strconv.Atoi(userIdStr)
	if err != nil {
		panic("Invalid user id")
	}

	var user app.User

	services.DeleteByIdUser(uint(photoId), &user, userReq.(app.User).ID)

	ctx.JSON(http.StatusOK, app.BaseResponse{
		Success: true,
		Message: "Successfully delete user",
	})
}
