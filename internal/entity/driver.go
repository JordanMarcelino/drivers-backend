package entity

import "github.com/shopspring/decimal"

type Driver struct {
	DriverCode            string
	Name                  string
	TotalPending          decimal.Decimal
	TotalConfirmed        decimal.Decimal
	TotalPaid             decimal.Decimal
	TotalAttendanceSalary decimal.Decimal
	TotalSalary           decimal.Decimal
	CountShipment         int
}
