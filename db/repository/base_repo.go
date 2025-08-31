package repository

import (
	"gorm.io/gorm"
)

type BasePostgresRepository[T any] struct {
	DB *gorm.DB
}

// Create new record
func (r *BasePostgresRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

// Get by id
func (r *BasePostgresRepository[T]) GetById(id uint, entity *T) error {
	return r.DB.First(entity, id).Error
}

// Get all
func (r *BasePostgresRepository[T]) GetAll(entities *[]T) error {
	return r.DB.Find(entities).Error
}

// Update
func (r *BasePostgresRepository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

// Delete
func (r *BasePostgresRepository[T]) Delete(entity *T) error {
	return r.DB.Delete(entity).Error
}
