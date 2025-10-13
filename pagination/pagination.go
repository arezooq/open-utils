package pagination

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

// NewFromRequest creates pagination params directly from a gin.Context
func NewFromRequest(c *gin.Context) *Params {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

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
func (p *Params) SetTotal(total int) {
	p.Total = total
	p.TotalPages = int(math.Ceil(float64(total) / float64(p.Limit)))
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
