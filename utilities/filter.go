package utilities

type Filter struct {
	Page   int `json:"page" query:"page"`
	Limit  int `json:"limit" query:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func (f *Filter) CalculateOffset() int {
	return f.Limit * (f.Page - 1)
}
