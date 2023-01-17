package utilities

import "math"

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type Meta struct {
	Limit     int `json:"limit"`
	Page      int `json:"page"`
	PerPage   int `json:"perPage"`
	PageCount int `json:"pageCount"`
	Total     int `json:"total"`
}

func BuildMeta(filter Filter, data int) Meta {
	pageCount := 0
	if filter.Limit > 0 {
		pageCount = int(math.Ceil(float64(filter.Total) / float64(filter.Limit)))
	}
	return Meta{
		Limit:     filter.Limit,
		Page:      filter.Page,
		PerPage:   data,
		PageCount: pageCount,
		Total:     filter.Total,
	}
}

func SetResponse(status string, message string, meta, result interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Meta:    meta,
		Result:  result,
	}
}
