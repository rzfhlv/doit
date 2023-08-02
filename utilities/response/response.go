package response

import (
	"math"

	"github.com/rzfhlv/doit/utilities/param"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type Meta struct {
	Limit     int   `json:"limit"`
	Page      int   `json:"page"`
	PerPage   int   `json:"perPage"`
	PageCount int   `json:"pageCount"`
	Total     int64 `json:"total"`
}

func BuildMeta(param param.Param, data int) Meta {
	pageCount := 0
	if param.Limit > 0 {
		pageCount = int(math.Ceil(float64(param.Total) / float64(param.Limit)))
	}
	return Meta{
		Limit:     param.Limit,
		Page:      param.Page,
		PerPage:   data,
		PageCount: pageCount,
		Total:     param.Total,
	}
}

func Set(status string, message string, meta, result interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Meta:    meta,
		Result:  result,
	}
}
