package repositories

import (
	"go-assignment-bootcamp/internal/app/models"
	"gorm.io/gorm"
)

type OrderDestroyRepository struct {
	gorm *gorm.DB
}

func NewOrderDestroyRepository(gorm *gorm.DB) *OrderDestroyRepository {
	return &OrderDestroyRepository{gorm}
}

func (repository *OrderDestroyRepository) Make(order *models.Order) error {
	result := repository.gorm.Delete(order)
	return result.Error
}
