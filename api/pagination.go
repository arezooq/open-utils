package api

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationParams struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Offset     int   `json:"offset"`
	Total      int64 `json:"total,omitempty"`
	TotalPages int64 `json:"total_pages,omitempty"`
}

func NewPaginationFromRequest(c *gin.Context) *PaginationParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	offset := (page - 1) * limit

	return &PaginationParams{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
}

func (p *PaginationParams) SetTotal(total int64) {
	p.Total = total
	p.TotalPages = int64(math.Ceil(float64(total) / float64(p.Limit)))
}

func (p *PaginationParams) JSON(data any) gin.H {
	return gin.H{
		"items":       data,
		"total":       p.Total,
		"page":        p.Page,
		"limit":       p.Limit,
		"total_pages": p.TotalPages,
	}
}
