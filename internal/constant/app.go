package constant

const (
	ResponseSuccessMessage = "success"
)

const (
	SHIPMENT_STATUS_RUNNING   = "RUNNING"
	SHIPMENT_STATUS_DONE      = "DONE"
	SHIPMENT_STATUS_CANCELLED = "CANCELLED"
)

const (
	SHIPMENT_COST_STATUS_PENDING   = "PENDING"
	SHIPMENT_COST_STATUS_CONFIRMED = "CONFIRMED"
	SHIPMENT_COST_STATUS_PAID      = "PAID"
)

var (
	SHIPMENT_COST_STATUSES = []string{SHIPMENT_COST_STATUS_PENDING, SHIPMENT_COST_STATUS_CONFIRMED, SHIPMENT_COST_STATUS_PAID}
)
