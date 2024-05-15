package adepters

import (
	"github.com/JimmyTanapon/Hexagonal/core"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) SaveOrder(order core.Order) error {
	if result := r.db.Create(&order); result != nil {
		return result.Error
	}
	return nil
}
