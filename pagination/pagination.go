package pagination

type Params struct {
	Page   int `json:"page"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func New(page, limit int) *Params {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit
	return &Params{Page: page, Limit: limit, Offset: offset}
}
