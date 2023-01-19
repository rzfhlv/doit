package utilities

type Param struct {
	Page   int   `json:"page" query:"page"`
	Limit  int   `json:"limit" query:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

func (f *Param) CalculateOffset() int {
	return f.Limit * (f.Page - 1)
}
