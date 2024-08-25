package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JordanMarcelino/drivers-backend/internal/constant"
	"github.com/JordanMarcelino/drivers-backend/internal/dto"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/apperror"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/utils/validationutils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var codeMap = map[int]int{
	apperror.DefaultClientErrorCode: http.StatusBadRequest,
	apperror.DefaultServerErrorCode: http.StatusInternalServerError,
	apperror.NotFoundErrorCode:      http.StatusNotFound,
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errLen := len(c.Errors)
		if errLen > 0 {
			err := c.Errors.Last()

			switch e := err.Err.(type) {
			case *json.SyntaxError:
				handleJsonSyntaxError(c)
			case *json.UnmarshalTypeError:
				handleJsonUnmarshalTypeError(c, e)
			case validator.ValidationErrors:
				handleValidationError(c, e)
			case *apperror.AppError:
				c.AbortWithStatusJSON(codeMap[e.GetCode()], dto.WebResponse[any]{
					Message: e.DisplayMessage(),
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.WebResponse[any]{
					Message: constant.InternalServerErrorMessage,
				})
			}
		}
	}
}

func handleJsonSyntaxError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: constant.JsonSyntaxErrorMessage,
	})
}

func handleJsonUnmarshalTypeError(c *gin.Context, err *json.UnmarshalTypeError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: fmt.Sprintf(constant.InvalidJsonValueTypeErrorMessage, err.Field),
	})
}

func handleValidationError(c *gin.Context, err validator.ValidationErrors) {
	ve := []dto.FieldError{}

	for _, fe := range err {
		ve = append(ve, dto.FieldError{
			Field:   fe.Field(),
			Message: validationutils.TagToMsg(fe),
		})
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: constant.ValidationErrorMessage,
		Errors:  ve,
	})
}
