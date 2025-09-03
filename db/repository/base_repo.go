package repository

import (
	"gorm.io/gorm"
)

type BasePostgresRepository[T any] struct {
	DB *gorm.DB
}

// Create new record
func (r *BasePostgresRepository[T]) Create(entity *T) (*T, error) {
	result := r.DB.Create(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

// Get by id
func (r *BasePostgresRepository[T]) GetById(id uint) (*T, error) {
	var entity T
	result := r.DB.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

// Get all
func (r *BasePostgresRepository[T]) GetAll() ([]T, error) {
	var entities []T
	result := r.DB.Find(entities)
	if result.Error != nil {
		return nil, result.Error
	}
	return entities, nil
}

// Update
func (r *BasePostgresRepository[T]) Update(entity *T) (*T, error) {
	result := r.DB.Save(entity)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}

// Delete
func (r *BasePostgresRepository[T]) Delete(entity *T) error {
	result := r.DB.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
