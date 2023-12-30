package helpers

import (
	"net/http"
	"photo-api/app"

	"github.com/gin-gonic/gin"
)

func BadRequestError(message string) {
	panic(app.ErrorCodeResponse{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	})
}

func UnauthorizedError(message string) {
	panic(app.ErrorCodeResponse{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	})
}

func NotFoundError(message string) {
	panic(app.ErrorCodeResponse{
		StatusCode: http.StatusNotFound,
		Message:    message,
	})
}

func ForbiddenError(message string) {
	panic(app.ErrorCodeResponse{
		StatusCode: http.StatusForbidden,
		Message:    message,
	})
}

func ServerErrorError(message string) {
	panic(app.ErrorCodeResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	})
}

func RecoverErrorResponse(ctx *gin.Context) {
	errorValue := recover()

	if errorValue == nil {
		return
	}

	switch value := errorValue.(type) {
	case app.ErrorCodeResponse:
		ctx.JSON(value.StatusCode, app.BaseResponse{
			Success: false,
			Message: value.Message,
		})
	case error:
		ctx.JSON(http.StatusInternalServerError, app.BaseResponse{
			Success: false,
			Message: value.Error(),
		})
	default:
		ctx.JSON(http.StatusInternalServerError, app.BaseResponse{
			Success: false,
			Message: value.(string),
		})
	}

	ctx.Abort()
}
