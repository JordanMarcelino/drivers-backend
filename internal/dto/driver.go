package dto

import (
	"github.com/JordanMarcelino/drivers-backend/internal/entity"
	"github.com/shopspring/decimal"
)

type DriverDTO struct {
	DriverCode            string          `json:"driver_code"`
	Name                  string          `json:"name"`
	TotalPending          decimal.Decimal `json:"total_pending"`
	TotalConfirmed        decimal.Decimal `json:"total_confirmed"`
	TotalPaid             decimal.Decimal `json:"total_paid"`
	TotalAttendanceSalary decimal.Decimal `json:"total_attendance_salary"`
	TotalSalary           decimal.Decimal `json:"total_salary"`
	CountShipment         int             `json:"count_shipment"`
}

type DriverQueryDTO struct {
	DriverCode string `form:"driver_code"`
	Status     string `form:"status" binding:"omitempty,shipment_cost_statuses"`
	Name       string `form:"name"`
	Month      int    `form:"month" binding:"required,numeric,gte=1,lte=12"`
	Year       int    `form:"year" binding:"required,numeric,gte=2010"`
	PageSize   int    `form:"page_size" binding:"omitempty,gte=1,lte=20"`
	Current    int    `form:"current" binding:"omitempty,min=1"`
}

func ConvertToDriverDTOs(drivers []*entity.Driver) []*DriverDTO {
	dtos := []*DriverDTO{}
	for _, driver := range drivers {
		dtos = append(dtos, ConvertToDriverDTO(driver))
	}

	return dtos
}

func ConvertToDriverDTO(driver *entity.Driver) *DriverDTO {
	return &DriverDTO{
		DriverCode:            driver.DriverCode,
		Name:                  driver.Name,
		TotalPending:          driver.TotalPending,
		TotalConfirmed:        driver.TotalConfirmed,
		TotalPaid:             driver.TotalPaid,
		TotalAttendanceSalary: driver.TotalAttendanceSalary,
		TotalSalary:           driver.TotalSalary,
		CountShipment:         driver.CountShipment,
	}
}
