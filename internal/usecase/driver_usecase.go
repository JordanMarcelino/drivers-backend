package usecase

import (
	"context"

	"github.com/JordanMarcelino/drivers-backend/internal/dto"
	"github.com/JordanMarcelino/drivers-backend/internal/entity"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/apperror"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/utils/pageutils"
	"github.com/JordanMarcelino/drivers-backend/internal/repository"
)

type DriverUseCase interface {
	FindAll(ctx context.Context, params *dto.DriverQueryDTO) ([]*dto.DriverDTO, int, error)
}

type driverUseCaseImpl struct {
	driverRepo repository.DriverRepository
}

func NewDriverUseCase(driverRepo repository.DriverRepository) *driverUseCaseImpl {
	return &driverUseCaseImpl{
		driverRepo: driverRepo,
	}
}

func (u driverUseCaseImpl) FindAll(ctx context.Context, params *dto.DriverQueryDTO) ([]*dto.DriverDTO, int, error) {
	pageutils.CreateDefaultParams(params)

	drivers, err := u.driverRepo.FindAll(ctx, params)
	if err != nil {
		return nil, 0, apperror.NewServerError(err)
	}

	page := params.Current
	limit := params.PageSize
	totalRows := len(drivers)

	switch {
	case totalRows > page*limit:
		drivers = drivers[limit*(page-1) : limit*page]
	case totalRows > (page-1)*limit:
		drivers = drivers[limit*(page-1):]
	default:
		drivers = []*entity.Driver{}
	}

	return dto.ConvertToDriverDTOs(drivers), totalRows, nil
}
