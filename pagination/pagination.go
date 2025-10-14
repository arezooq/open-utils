package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	Offset     int64 `json:"offset"`
	Total      int64 `json:"total,omitempty"`
	TotalPages int64 `json:"total_pages,omitempty"`
}

// NewFromRequest creates pagination params directly from a gin.Context
func NewFromRequest(c *gin.Context) *Params {
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)

	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	return &Params{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

// SetTotal sets total count and calculates total pages
func (p *Params) SetTotal(total int64) {
	p.Total = total
	p.TotalPages = total / p.Limit
}

// JSON formats the pagination result
func (p *Params) JSON(data any) gin.H {
	return gin.H{
		"items":       data,
		"total":       p.Total,
		"page":        p.Page,
		"limit":       p.Limit,
		"total_pages": p.TotalPages,
	}
}
