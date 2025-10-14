package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Filter struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type Search struct {
	Field string `json:"field"`
	Query string `json:"query"`
}

type Order struct {
	Field     string `json:"field"`
	Direction string `json:"direction"` // asc | desc
}

type QueryParams struct {
	Filters []Filter `json:"filters"`
	Search  []Search `json:"search"`
	Orders  []Order  `json:"orders"`
}

func NewQueryFromRequest(c *gin.Context) *QueryParams {
	var params QueryParams

	if err := c.ShouldBindJSON(&params); err == nil && (len(params.Filters) > 0 || len(params.Search) > 0) {
		return &params
	}

	if raw := c.Query("filters"); raw != "" {
		_ = json.Unmarshal([]byte(raw), &params.Filters)
	}
	if raw := c.Query("search"); raw != "" {
		_ = json.Unmarshal([]byte(raw), &params.Search)
	}
	if raw := c.Query("orders"); raw != "" {
		_ = json.Unmarshal([]byte(raw), &params.Orders)
	}

	return &params
}

func ApplyFilters(db *gorm.DB, filters []Filter) *gorm.DB {
	for _, f := range filters {
		db = db.Where(f.Field+" = ?", f.Value)
	}
	return db
}

func ApplySearch(db *gorm.DB, searches []Search) *gorm.DB {
	for _, s := range searches {
		db = db.Where(s.Field+" LIKE ?", "%"+s.Query+"%")
	}
	return db
}

func ApplyOrder(db *gorm.DB, orders []Order) *gorm.DB {
	for _, o := range orders {
		dir := o.Direction
		if dir == "" {
			dir = "asc"
		}
		db = db.Order(o.Field + " " + dir)
	}
	return db
}
