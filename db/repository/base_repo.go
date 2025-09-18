package repository

import (
	"github.com/arezooq/open-utils/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BasePostgresRepository[T any] struct {
	DB *gorm.DB
}

func NewBasePostgresRepository[T any](db *gorm.DB) *BasePostgresRepository[T] {
	return &BasePostgresRepository[T]{DB: db}
}

// Create new record
func (r *BasePostgresRepository[T]) Create(entity *T) (*T, error) {
	result := r.DB.Create(entity)
	if result.Error != nil {
		return nil, errors.ErrInternal
	}
	return entity, nil
}

// Get by id
func (r *BasePostgresRepository[T]) GetById(id uint) (*T, error) {
	var entity T
	result := r.DB.First(&entity, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, errors.ErrInternal
	}
	return &entity, nil
}

// Get all
func (r *BasePostgresRepository[T]) GetAll() ([]T, error) {
	var entities []T
	result := r.DB.Find(&entities)
	if result.Error != nil {
		return nil, errors.ErrInternal
	}
	return entities, nil
}

// Update
func (r *BasePostgresRepository[T]) Update(id uuid.UUID, updates map[string]any) (*T, error) {
	var entity T
	result := r.DB.Model(&entity).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return nil, errors.ErrInternal
	}

	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, errors.ErrNotFound
	}
	return &entity, nil
}

// Delete
func (r *BasePostgresRepository[T]) Delete(id uint) error {
	var entity T
	result := r.DB.Delete(&entity, id)
	
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.ErrNotFound
		}
		return errors.ErrInternal
	}

	if result.RowsAffected == 0 {
		return errors.ErrNotFound
	}

	return nil
}
