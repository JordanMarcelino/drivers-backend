package apperror

import "github.com/JordanMarcelino/drivers-backend/internal/constant"

func NewServerError(err error) AppError {
	msg := constant.InternalServerErrorMessage

	return NewAppError(err, DefaultServerErrorCode, msg, nil)
}
