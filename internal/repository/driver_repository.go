package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/JordanMarcelino/drivers-backend/internal/dto"
	"github.com/JordanMarcelino/drivers-backend/internal/entity"
)

const (
	findAllQuery = `
	SELECT 
		d.driver_code,
		d.name,
		COALESCE(
			SUM(CASE WHEN sc.cost_status = 'PENDING' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END), 0
		) AS total_pending,
		COALESCE(
			SUM(CASE WHEN sc.cost_status = 'CONFIRMED' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END), 0
		) AS total_confirmed,
		COALESCE(
			SUM(CASE WHEN sc.cost_status = 'PAID' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END), 0
		) AS total_paid,
		COALESCE(
			COUNT(DISTINCT CASE WHEN da.attendance_status = TRUE THEN da.attendance_date END) * vc.value, 0
		) AS total_attendance_salary,
		COALESCE(
			SUM(CASE WHEN sc.cost_status IN ('PENDING', 'CONFIRMED', 'PAID') AND s.shipment_status != 'CANCELLED' THEN sc.	total_costs ELSE 0 END) + 
        	COUNT(DISTINCT CASE WHEN da.attendance_status = TRUE THEN da.attendance_date END) * vc.value, 0
		) AS total_salary,
		COALESCE(
			COUNT(DISTINCT CASE WHEN s.shipment_status != 'CANCELLED' THEN sc.shipment_no END), 0
		) AS count_shipment
	FROM
		drivers d
	LEFT JOIN
		driver_attendances da ON d.driver_code = da.driver_code AND da.attendance_status = TRUE
	LEFT JOIN
		shipment_costs sc ON d.driver_code = sc.driver_code
	LEFT JOIN
		shipments s ON sc.shipment_no = s.shipment_no
	LEFT JOIN
		variable_configs vc ON vc.key = 'DRIVER_MONTHLY_ATTENDANCE_SALARY'
	WHERE
		(EXTRACT(MONTH FROM s.shipment_date) = $1) AND
		(EXTRACT(YEAR FROM s.shipment_date) = $2) AND
		(d.driver_code = $3 OR $3 = '') AND
		(LOWER(d.name) LIKE LOWER('%' || $4 || '%') OR $4 = '')
	GROUP BY
		d.driver_code, d.name, vc.value
	HAVING
		($5 = '') OR	
		($5 = 'PENDING' AND SUM(CASE WHEN sc.cost_status = 'PENDING' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END) > 0) OR
		($5 = 'CONFIRMED' AND SUM(CASE WHEN sc.cost_status = 'CONFIRMED' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END) > 0) OR
		($5 = 'PAID' AND SUM(CASE WHEN sc.cost_status = 'PAID' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END) > 0 AND
		SUM(CASE WHEN sc.cost_status = 'CONFIRMED' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END) = 0 AND
		SUM(CASE WHEN sc.cost_status = 'PENDING' AND s.shipment_status != 'CANCELLED' THEN sc.total_costs ELSE 0 END) = 0)
	`
)

type DriverRepository interface {
	FindAll(ctx context.Context, params *dto.DriverQueryDTO) ([]*entity.Driver, error)
}

type driverRepositoryImpl struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *driverRepositoryImpl {
	return &driverRepositoryImpl{
		db: db,
	}
}

func (r *driverRepositoryImpl) FindAll(ctx context.Context, params *dto.DriverQueryDTO) ([]*entity.Driver, error) {
	rows, err := r.db.QueryContext(
		ctx,
		findAllQuery,
		params.Month,
		params.Year,
		params.DriverCode,
		params.Name,
		params.Status,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	drivers := []*entity.Driver{}
	for rows.Next() {
		driver := new(entity.Driver)
		if err := rows.Scan(
			&driver.DriverCode,
			&driver.Name,
			&driver.TotalPending,
			&driver.TotalConfirmed,
			&driver.TotalPaid,
			&driver.TotalAttendanceSalary,
			&driver.TotalSalary,
			&driver.CountShipment,
		); err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return drivers, nil
}
