package repository

import (
	"github.com/arezooq/open-utils/errors"
	"gorm.io/gorm"
)

type BasePostgresRepository[T any] struct {
	DB *gorm.DB
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
func (r *BasePostgresRepository[T]) GetById(id string) (*T, error) {
	var entity T
	result := r.DB.Where("id = ?", id).First(&entity)
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
func (r *BasePostgresRepository[T]) Update(id string, updates map[string]any) (*T, error) {
    var existing T
    if err := r.DB.First(&existing, "id = ?", id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, errors.ErrNotFound
        }
        return nil, errors.ErrInternal
    }

    result := r.DB.Model(&existing).Updates(updates)
    if result.Error != nil {
        return nil, errors.ErrInternal
    }

    return &existing, nil
}


// Delete
func (r *BasePostgresRepository[T]) Delete(id string) error {
	var entity T
	result := r.DB.Where("id = ?", id).Delete(&entity, id)
	
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
