package httphandler

import (
	"net/http"

	"github.com/JordanMarcelino/drivers-backend/internal/constant"
	"github.com/JordanMarcelino/drivers-backend/internal/dto"
	"github.com/JordanMarcelino/drivers-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type DriverHandler struct {
	driverUseCase usecase.DriverUseCase
}

func NewDriverHandler(driverUseCase usecase.DriverUseCase) *DriverHandler {
	return &DriverHandler{
		driverUseCase: driverUseCase,
	}
}

func (h *DriverHandler) FindAll(ctx *gin.Context) {
	params := new(dto.DriverQueryDTO)
	if err := ctx.ShouldBindQuery(params); err != nil {
		ctx.Error(err)
		return
	}

	drivers, totalRow, err := h.driverUseCase.FindAll(ctx, params)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, dto.PageResponse[[]*dto.DriverDTO]{
		Message:  constant.ResponseSuccessMessage,
		Data:     drivers,
		TotalRow: totalRow,
		Current:  params.Current,
		PageSize: params.PageSize,
	})
}
