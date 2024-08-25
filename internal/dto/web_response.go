package dto

type WebResponse[T any] struct {
	Message string       `json:"message"`
	Data    T            `json:"data,omitempty"`
	Errors  []FieldError `json:"errors,omitempty"`
}

type PageResponse[T any] struct {
	Message  string `json:"message"`
	Data     T      `json:"data"`
	TotalRow int    `json:"total_row"`
	Current  int    `json:"current"`
	PageSize int    `json:"page_size"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
