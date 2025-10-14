package api

import "gorm.io/gorm"

type Filter struct {
	Field string
	Value any
}

type Search struct {
	Field string
	Query string
}

func ApplyFilters[T any](db *gorm.DB, filters []Filter, searches []Search) *gorm.DB {
	for _, f := range filters {
		db = db.Where(f.Field+" = ?", f.Value)
	}

	for _, s := range searches {
		db = db.Where(s.Field+" LIKE ?", "%"+s.Query+"%")
	}

	return db
}
