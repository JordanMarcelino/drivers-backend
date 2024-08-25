package validationutils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/JordanMarcelino/drivers-backend/internal/constant"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func TagToMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "len":
		return fmt.Sprintf("%s must be %v characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must not exceed %v characters", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be lower than or equal to %v", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be %v characters long", fe.Field(), fe.Param())
	case "numeric":
		return fmt.Sprintf("%s must be numeric", fe.Field())
	case "shipment_cost_statuses":
		return fmt.Sprintf("%s possible values are ['PENDING', 'CONFIRMED', 'PAID']", fe.Field())
	default:
		return "invalid input"
	}
}

func FormTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}

func DecimalType(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(decimal.Decimal); ok {
		return valuer.String()
	}
	return nil
}

func ShipmentCostStatusValidator(fl validator.FieldLevel) bool {
	status, ok := fl.Field().Interface().(string)
	if ok {
		for _, STATUS := range constant.SHIPMENT_COST_STATUSES {
			if status == STATUS {
				return true
			}
		}
		return false
	}

	return false
}
