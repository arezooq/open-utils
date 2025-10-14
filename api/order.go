package api

import "gorm.io/gorm"

type Order struct {
	Field     string
	Direction string // "asc" || "desc"
}

func ApplyOrder[T any](db *gorm.DB, orders []Order) *gorm.DB {
	for _, o := range orders {
		db = db.Order(o.Field + " " + o.Direction)
	}
	return db
}
