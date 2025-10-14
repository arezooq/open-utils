package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Filter struct {
	Field string
	Value any
}

type Search struct {
	Field string
	Query string
}

type Order struct {
	Field     string
	Direction string // asc | desc
}

type QueryParams struct {
	Search  string
	Filters map[string]interface{}
	Orders  []string
}

func NewQueryFromRequest(c *gin.Context) *QueryParams {
	search := c.Query("search")
	orders := c.QueryArray("order")
	filters := map[string]interface{}{}

	for key, vals := range c.Request.URL.Query() {
		if key != "search" && key != "order" && key != "page" && key != "limit" {
			filters[key] = vals[0]
		}
	}

	return &QueryParams{
		Search:  search,
		Orders:  orders,
		Filters: filters,
	}
}

func ApplyFilters(db *gorm.DB, filters []Filter, searches []Search) *gorm.DB {
	for _, f := range filters {
		db = db.Where(f.Field+" = ?", f.Value)
	}
	for _, s := range searches {
		db = db.Where(s.Field+" LIKE ?", "%"+s.Query+"%")
	}
	return db
}

func ApplyOrder(db *gorm.DB, orders []Order) *gorm.DB {
	for _, o := range orders {
		db = db.Order(o.Field + " " + o.Direction)
	}
	return db
}
