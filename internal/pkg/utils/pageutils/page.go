package pageutils

import (
	"github.com/JordanMarcelino/drivers-backend/internal/dto"
)

var (
	defaultPageSize = 10
	defaultCurrent  = 1
)

func CreateDefaultParams(params *dto.DriverQueryDTO) {
	if params.PageSize == 0 {
		params.PageSize = defaultPageSize
	}

	if params.Current == 0 {
		params.Current = defaultCurrent
	}
}
